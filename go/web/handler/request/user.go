package request

type UserGet struct {
	IdToken string `json:"Authorization"`
}

type UserPut struct {
	IdToken string `json:"Authorization"`
	Name    string `json:"name"`
}

type UserDelete struct {
	IdToken string `json:"Authorization"`
}
