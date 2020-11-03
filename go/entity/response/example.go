package response

import "time"

type Example struct {
	ID        int        `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
}

type ExampleGet []Example

type ExampleGetById Example

type ExamplePost Example

type ExamplePutById Example

type ExampleDeleteById struct{}
