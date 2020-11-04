package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Plant は植物の永続化を担当するリポジトリです。
type Plant interface {
	FindByID(id string) (*entity.Plant, error)
	FindAll() (*entity.Plants, error)
	Store(plant *entity.Plant)  (*entity.Plant, error)
	UpdateByID(id string, plant *entity.Plant) (*entity.Plant, error)
	DeleteByID(id string) (*entity.Plant, error)
}