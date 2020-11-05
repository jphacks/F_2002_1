package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Watering は水やりの記録の永続化を担当するリポジトリです。
type Watering interface {
	FindByID(id int) (*entity.Watering, error)
	FindAll() (*entity.Waterings, error)
	Store(watering *entity.Watering) (*entity.Watering, error)
	UpdateByID(watering *entity.Watering) (*entity.Watering, error)
	DeleteByID(id int) (*entity.Watering, error)
}
