package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Watering = &WateringRepository{}

// WateringRepository は repository.WateringRepository を満たす構造体です。
type WateringRepository struct {
	db *gorm.DB
}

// NewWateringRepository はWateringRepositoryのポインタを生成する関数です。
func NewWateringRepository(db *gorm.DB) *WateringRepository {
	return &WateringRepository{db: db}
}

// FindByID は指定されたIDを持つ水やりの記録を取得します。
func (r *WateringRepository) FindByID(id int) (*entity.Watering, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つ水やりの記録を取得します。
func (r *WateringRepository) FindAll() (*entity.Waterings, error) {
	return nil, nil
}

// Store は水やりの記録を新規保存します。
func (r *WateringRepository) Store(watering *entity.Watering) (*entity.Watering, error) {
	return nil, nil
}

// UpdateByID は水やりの記録の情報を更新します。
func (r *WateringRepository) UpdateByID(watering *entity.Watering) (*entity.Watering, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つ水やりの記録を削除します。
func (r *WateringRepository) DeleteByID(id int) (*entity.Watering, error) {
	return nil, nil
}
