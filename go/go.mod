module github.com/jphacks/F_2002_1/go

go 1.15

require (
	cloud.google.com/go/firestore v1.3.0 // indirect
	cloud.google.com/go/storage v1.12.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo/v4 v4.1.17
	google.golang.org/api v0.34.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0 // Need to use github.com/oxequa/realize, used in ./Dockerfile
