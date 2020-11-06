package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Temperature = &TemperatureRepository{}

// TemperatureRepository は repository.TemperatureRepository を満たす構造体です。
type TemperatureRepository struct {
	db *gorm.DB
}

// NewTemperatureRepository はTemperatureRepositoryのポインタを生成する関数です。
func NewTemperatureRepository(db *gorm.DB) *TemperatureRepository {
	return &TemperatureRepository{db: db}
}

// FindByID は指定されたIDを持つ気温を取得します。
func (r *TemperatureRepository) FindByID(id int) (*entity.Temperature, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つ気温を取得します。
func (r *TemperatureRepository) FindAll() (*entity.Temperatures, error) {
	return nil, nil
}

// Store は気温を新規保存します。
func (r *TemperatureRepository) Store(temperature *entity.Temperature) (*entity.Temperature, error) {
	return nil, nil
}

// UpdateByID は気温の情報を更新します。
func (r *TemperatureRepository) UpdateByID(temperature *entity.Temperature) (*entity.Temperature, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つ気温を削除します。
func (r *TemperatureRepository) DeleteByID(id int) error {
	return nil
}
