package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/irodavlas/api-gateway/model"
	"github.com/irodavlas/api-gateway/utils"

	"github.com/golang-jwt/jwt/v4"

	response "github.com/irodavlas/common-response"

	"github.com/labstack/echo/v4"
)

const (
	users_service_url = "http://users-service:5000"
)

// define auth methods interface for the routes
type IAuthService interface {
	Login(c echo.Context) error
	ValidateToken(c echo.Context) error
	SignIn(c echo.Context) error
}
type AuthService struct {
	Sender    response.Sender
	SecretKey string
}

func NewAuthService(secretKey string) *AuthService {
	return &AuthService{
		SecretKey: secretKey,
		Sender:    *response.NewSender(),
	}
}

func (auth *AuthService) ValidateToken(c echo.Context) error {
	log.Println("validating")
	return nil

}
func (auth *AuthService) SignIn(c echo.Context) error {
	//check if user exists
	//create user -> done

	resp, err := utils.RequestMicroservice("POST", users_service_url+"/create", c.Request().Body)
	if err != nil {
		return auth.Sender.Error(c, http.StatusInternalServerError, "Error requesting USER microservice", err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return auth.Sender.Error(c, http.StatusInternalServerError, "Error reading body of response from USER microservice", err.Error())
	}
	microserviceResponse := model.ResponseGetUser{}
	if err := json.Unmarshal(body, &microserviceResponse); err != nil {
		return auth.Sender.Error(c, http.StatusInternalServerError, "invalid response json", err.Error())
	}
	if resp.StatusCode == http.StatusOK {
		token, err := auth.generateToken()
		if err != nil {
			return auth.Sender.Error(c, http.StatusInternalServerError, "Error generating JWT token", err.Error())
		}
		return auth.Sender.SuccessWithToken(c, http.StatusOK, "Success", microserviceResponse.Data, token)
	}
	return auth.Sender.Error(c, resp.StatusCode, "Error serving request", "Error occurred while signing in the user")
}
func (auth *AuthService) Login(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return auth.Sender.Error(c, http.StatusBadRequest, "Error while binding response body to User struct", err.Error())
	}
	url := users_service_url + fmt.Sprintf("/get?username=%s", user.Username)
	resp, err := utils.RequestMicroservice("GET", url, nil)
	if err != nil {
		return auth.Sender.Error(c, http.StatusInternalServerError, "Error making request to User service", err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	microserviceResponse := model.ResponseGetUser{}
	if err := json.Unmarshal(body, &microserviceResponse); err != nil {
		return auth.Sender.Error(c, http.StatusInternalServerError, "invalid response json", err.Error())
	}
	//still not sure how to implement an admin view, when fronted will be done ill get back to this
	if resp.StatusCode == http.StatusOK {
		if !verifyPassword(user.Password, microserviceResponse.Data.Password) {
			return auth.Sender.Error(c, http.StatusBadRequest, "Password or Username are wrong", "password did not match")
		}
		token, err := auth.generateToken()
		if err != nil {

			return auth.Sender.Error(c, http.StatusInternalServerError, "Error generating JWT token", err.Error())
		}
		return auth.Sender.SuccessWithToken(c, resp.StatusCode, microserviceResponse.Message, microserviceResponse.Data, token)
	}

	return auth.Sender.Error(c, resp.StatusCode, microserviceResponse.Message, "")

}

func (auth *AuthService) generateToken() (string, error) {
	now := time.Now()
	expirationDate := now.Add(1 * time.Hour)
	claims := &jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expirationDate),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(auth.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// consider implementing some real validation or switch to google auth
func verifyPassword(reqPassword, dbPassword string) bool {
	return reqPassword == dbPassword
}

func HandleIndexPage(c echo.Context) error {
	templ, _ := template.New("").ParseFiles("frontend/templates/layout.html")
	return templ.ExecuteTemplate(c.Response().Writer, "base", nil)
}
