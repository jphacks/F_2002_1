package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Water = &WaterRepository{}

// WaterRepository は repository.WaterRepository を満たす構造体です。
type WaterRepository struct {
	db *gorm.DB
}

// NewWaterRepository はWaterRepositoryのポインタを生成する関数です。
func NewWaterRepository(db *gorm.DB) *WaterRepository {
	return &WaterRepository{db: db}
}

// FindByID は指定されたIDを持つ水分量を取得します。
func (r *WaterRepository) FindByID(id int) (*entity.Water, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つ水分量を取得します。
func (r *WaterRepository) FindAll() (*entity.Waters, error) {
	return nil, nil
}

// Store は水分量を新規保存します。
func (r *WaterRepository) Store(water *entity.Water) (*entity.Water, error) {
	return nil, nil
}

// UpdateByID は水分量の情報を更新します。
func (r *WaterRepository) UpdateByID(water *entity.Water) (*entity.Water, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つ水分量を削除します。
func (r *WaterRepository) DeleteByID(id int) error {
	return nil
}
