module github.com/jphacks/F_2002_1/go

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo/v4 v4.1.17
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0 // Need to use github.com/oxequa/realize, used in ./Dockerfile
