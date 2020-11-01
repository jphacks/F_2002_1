package usecase

import "github.com/jphacks/F_2002_1/go/entity/model"

type ExampleRepository interface {
	FindById(id int) (model.Example, error)
	FindAll() (model.Examples, error)
	Store(model.Example) (model.Example, error)
	Update(model.Example) (model.Example, error)
	DeleteById(model.Example) error
}
