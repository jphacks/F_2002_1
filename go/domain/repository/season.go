package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// Season は季節の永続化を担当するリポジトリです。
type Season interface {
	FindByID(id int) (*entity.Season, error)
	FindAll() (*entity.Seasons, error)
	Store(season *entity.Season) (*entity.Season, error)
	UpdateByID(season *entity.Season) (*entity.Season, error)
	DeleteByID(id int) error
}
