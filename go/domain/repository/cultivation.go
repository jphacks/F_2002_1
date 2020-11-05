package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Cultivation は栽培している植物の永続化を担当するリポジトリです。
type Cultivation interface {
	FindByID(id int) (*entity.Cultivation, error)
	FindAllByUID(uid int) (*entity.Cultivations, error)
	CheckByIDUID(id int, uid int) (bool, error)
	Store(cultivation *entity.Cultivation) (*entity.Cultivation, error)
	UpdateByID(cultivation *entity.Cultivation) (*entity.Cultivation, error)
	DeleteByID(id int) (*entity.Cultivation, error)
	CheckByIDUID(id int, uid string) (bool, error)
}
