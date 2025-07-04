package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	response "github.com/irodavlas/common-response"
	"github.com/labstack/echo/v4"
)

func router_proxy(url string, sender response.Sender) echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			sender.Error(c, http.StatusInternalServerError, "Error reading request body", err.Error())
		}
		client := &http.Client{}
		req, _ := http.NewRequest(c.Request().Method, url+c.Request().URL.Path, bytes.NewReader(body))
		req.Header.Set("Authorization", c.Request().Header.Get("Authorization"))
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			sender.Error(c, http.StatusInternalServerError, "Service unavailable", err.Error())
		}
		defer resp.Body.Close()

		body_, err := io.ReadAll(resp.Body)
		if err != nil {
			sender.Error(c, http.StatusInternalServerError, "Error handling server response", err.Error())
		}
		var Json interface{}
		if err := json.Unmarshal(body_, &Json); err != nil {
			return c.JSON(resp.StatusCode, string(body_))
		}
		return c.JSON(resp.StatusCode, Json)
	}
}
