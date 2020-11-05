package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Cultivation は栽培している植物の永続化を担当するリポジトリです。
type Cultivation interface {
	FindByID(id int) (*entity.Cultivation, error)
	FindAll() (*entity.Cultivations, error)
	Store(cultivation *entity.Cultivation)  (*entity.Cultivation, error)
	UpdateByID(id int, cultivation *entity.Cultivation) (*entity.Cultivation, error)
	DeleteByID(id int) (*entity.Cultivation, error)
}
