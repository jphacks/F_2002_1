package response

import "github.com/jphacks/F_2002_1/go/web/handler/openapi"

type UsersGet []openapi.User

type UsersGetById openapi.User

type UsersPost openapi.User

type UsersPutById openapi.User

type UsersDeleteById struct{}
