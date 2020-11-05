package request

type UsersGet struct{}

type UsersGetByID struct {
	UserID int `param:"user_id"`
}

type UsersPost struct {
	Name string `json:"name"`
}

type UsersPutByID struct {
	UserID int    `param:"user_id"`
	Name   string `json:"name"`
}

type UsersDeleteByID struct {
	UserID int `param:"user_id"`
}
