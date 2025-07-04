module irodavlas/website

go 1.21.3

require (
	github.com/joho/godotenv v1.5.1
	github.com/labstack/echo-jwt v0.0.0-20221127215225-c84d41a71003
	github.com/labstack/echo/v4 v4.13.3
)

replace github.com/irodavlas/common-response => ./common/response/

require (
	github.com/golang-jwt/jwt/v4 v4.5.1 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.8.0 // indirect
)
