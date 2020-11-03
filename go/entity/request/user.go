package request

type UserGet struct{}

type UserPut struct {
	Name string `json:"name"`
}

type UserDelete struct{}
