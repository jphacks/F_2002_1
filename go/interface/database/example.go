package database

import "github.com/jphacks/F_2002_1/go/entity/model"

type ExampleRepository struct {
	SqlHandler
}

func (repo *ExampleRepository) FindById(id int) (example model.Example, err error) {
	if err = repo.Set("gorm:auto_preload", true).First(&example, id).Error; err != nil {
		return
	}
	return
}

func (repo *ExampleRepository) FindAll() (examples model.Examples, err error) {
	if err = repo.Set("gorm:auto_preload", true).Find(&examples).Error; err != nil {
		return
	}
	return
}

func (repo *ExampleRepository) Store(ex model.Example) (example model.Example, err error) {
	if err = repo.Set("gorm:auto_preload", true).Create(&ex).Error; err != nil {
		return
	}
	example = ex
	return
}

func (repo *ExampleRepository) Update(ex model.Example) (example model.Example, err error) {
	if err = repo.Set("gorm:auto_preload", true).Model(&model.Example{}).Update(&ex).First(&ex).Error; err != nil {
		return
	}
	example = ex
	return
}

func (repo *ExampleRepository) DeleteById(ex model.Example) (err error) {
	if err = repo.Delete(&ex).Error; err != nil {
		return
	}
	return
}
