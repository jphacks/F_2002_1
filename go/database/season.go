package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Season = &SeasonRepository{}

// SeasonRepository は repository.SeasonRepository を満たす構造体です。
type SeasonRepository struct {
	db *gorm.DB
}

// NewSeasonRepository はSeasonRepositoryのポインタを生成する関数です。
func NewSeasonRepository(db *gorm.DB) *SeasonRepository {
	return &SeasonRepository{db: db}
}

// FindByID は指定されたIDを持つ季節を取得します。
func (r *SeasonRepository) FindByID(id int) (*entity.Season, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つ季節を取得します。
func (r *SeasonRepository) FindAll() (*entity.Seasons, error) {
	return nil, nil
}

// Store は季節を新規保存します。
func (r *SeasonRepository) Store(season *entity.Season) (*entity.Season, error) {
	return nil, nil
}

// UpdateByID は季節の情報を更新します。
func (r *SeasonRepository) UpdateByID(season *entity.Season) (*entity.Season, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つ季節を削除します。
func (r *SeasonRepository) DeleteByID(id int) (*entity.Season, error) {
	return nil, nil
}
