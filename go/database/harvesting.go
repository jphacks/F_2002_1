package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Harvesting = &HarvestingRepository{}

// HarvestingRepository は repository.HarvestingRepository を満たす構造体です。
type HarvestingRepository struct {
	db *gorm.DB
}

// NewHarvestingRepository はHarvestingRepositoryのポインタを生成する関数です。
func NewHarvestingRepository(db *gorm.DB) *HarvestingRepository {
	return &HarvestingRepository{db: db}
}

// FindByID は指定されたIDを持つ収穫の記録を取得します。
func (r *HarvestingRepository) FindByID(id string) (*entity.Harvesting, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つ収穫の記録を取得します。
func (r *HarvestingRepository) FindAll() (*entity.Harvestings, error) {
	return nil, nil
}
// Store は収穫の記録を新規保存します。
func (r *HarvestingRepository) Store(harvesting *entity.Harvesting) (*entity.Harvesting, error) {
	return nil, nil
}

// UpdateByID は収穫の記録の情報を更新します。
func (r *HarvestingRepository) UpdateByID(id string, harvesting *entity.Harvesting) (*entity.Harvesting, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つ収穫の記録を削除します。
func (r *HarvestingRepository) DeleteByID(id string) (*entity.Harvesting, error) {
	return nil, nil
}
