package request

type UserGet struct{}

type UserGetById struct {
	UserId int `json:"user_id"`
}

type UserPost struct {
	Name string `json:"name"`
}

type UserPutById struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type UserDeleteById struct {
	UserId int `json:"user_id"`
}
