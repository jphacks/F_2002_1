package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Plant = &PlantRepository{}

// PlantRepository は repository.PlantRepository を満たす構造体です。
type PlantRepository struct {
	db *gorm.DB
}

// NewPlantRepository はPlantRepositoryのポインタを生成する関数です。
func NewPlantRepository(db *gorm.DB) *PlantRepository {
	return &PlantRepository{db: db}
}

// FindByID は指定されたIDを持つ植物を取得します。
func (r *PlantRepository) FindByID(id int) (*entity.Plant, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つ植物を取得します。
func (r *PlantRepository) FindAll() (*entity.Plants, error) {
	return nil, nil
}
// Store は植物を新規保存します。
func (r *PlantRepository) Store(plant *entity.Plant) (*entity.Plant, error) {
	return nil, nil
}

// UpdateByID は植物の情報を更新します。
func (r *PlantRepository) UpdateByID(id int, plant *entity.Plant) (*entity.Plant, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つ植物を削除します。
func (r *PlantRepository) DeleteByID(id int) (*entity.Plant, error) {
	return nil, nil
}
