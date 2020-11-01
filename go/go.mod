module github.com/jphacks/F_2002_1/go

go 1.15

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo/v4 v4.1.17
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0 // Need to use github.com/oxequa/realize, used in ./Dockerfile
