package usecase

import "github.com/jphacks/F_2002_1/go/entity/model"

type ExampleInteractor struct {
	ExampleRepository ExampleRepository
}

func (interactor *ExampleInteractor) ExampleById(id int) (example model.Example, err error) {
	example, err = interactor.ExampleRepository.FindById(id)
	return
}

func (interactor *ExampleInteractor) Examples() (examples model.Examples, err error) {
	examples, err = interactor.ExampleRepository.FindAll()
	return
}

func (interactor *ExampleInteractor) Add(ex model.Example) (example model.Example, err error) {
	example, err = interactor.ExampleRepository.Store(ex)
	return
}

func (interactor *ExampleInteractor) Update(ex model.Example) (example model.Example, err error) {
	example, err = interactor.ExampleRepository.Update(ex)
	return
}

func (interactor *ExampleInteractor) DeleteById(ex model.Example) (err error) {
	err = interactor.ExampleRepository.DeleteById(ex)
	return
}
