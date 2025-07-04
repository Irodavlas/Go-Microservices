module github.com/irodavlas/api-gateway

go 1.23.0

toolchain go1.23.9

require (
	github.com/a-h/templ v0.3.865
	github.com/golang-jwt/jwt/v4 v4.5.1
	github.com/irodavlas/common-response v0.0.0-00010101000000-000000000000
	github.com/labstack/echo-jwt v0.0.0-20221127215225-c84d41a71003
	github.com/labstack/echo/v4 v4.13.3
)

replace github.com/irodavlas/common-response => ../common/response

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	golang.org/x/time v0.8.0 // indirect
)
