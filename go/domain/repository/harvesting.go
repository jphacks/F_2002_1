package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Harvesting は収穫の記録の永続化を担当するリポジトリです。
type Harvesting interface {
	FindByID(id int) (*entity.Harvesting, error)
	FindAll() (*entity.Harvestings, error)
	Store(harvesting *entity.Harvesting) (*entity.Harvesting, error)
	UpdateByID(harvesting *entity.Harvesting) (*entity.Harvesting, error)
	DeleteByID(id int) error
}
