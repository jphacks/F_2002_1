package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Water は水分量の永続化を担当するリポジトリです。
type Water interface {
	FindByID(id int) (*entity.Water, error)
	FindAll() (*entity.Waters, error)
	Store(water *entity.Water) (*entity.Water, error)
	UpdateByID(water *entity.Water) (*entity.Water, error)
	DeleteByID(id int) error
}
