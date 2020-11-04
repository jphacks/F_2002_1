package request

type ExampleGet struct{}

type ExampleGetById struct {
	ExampleId int `json:"example_id"`
}

type ExamplePost struct {
	Name string `json:"name"`
}

type ExamplePutById struct {
	ExampleId int    `json:"example_id"`
	Name      string `json:"name"`
}

type ExampleDeleteById struct {
	ExampleId int `json:"example_id"`
}
