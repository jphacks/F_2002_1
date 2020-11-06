package request

type UsersGet struct{}

type UsersGetByID struct {
	UserID int `param:"id"`
}

type UsersPost struct {
	Name string `json:"name"`
}

type UsersPutByID struct {
	UserID int    `param:"id"`
	Name   string `json:"name"`
}

type UsersDeleteByID struct {
	UserID int `param:"id"`
}
