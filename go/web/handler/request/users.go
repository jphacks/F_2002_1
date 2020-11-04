package request

type UsersGet struct{}

type UsersGetById struct {
	UserId int `json:"user_id"`
}

type UsersPost struct {
	Name string `json:"name"`
}

type UsersPutById struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type UsersDeleteById struct {
	UserId int `json:"user_id"`
}
