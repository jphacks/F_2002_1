package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.Cultivation = &CultivationRepository{}

// CultivationRepository は repository.CultivationRepository を満たす構造体です。
type CultivationRepository struct {
	db *gorm.DB
}

// NewCultivationRepository はCultivationRepositoryのポインタを生成する関数です。
func NewCultivationRepository(db *gorm.DB) *CultivationRepository {
	return &CultivationRepository{db: db}
}

// FindByID は指定されたIDを持つ栽培している植物を取得します。
func (r *CultivationRepository) FindByID(id int) (*entity.Cultivation, error) {
	var cultivation entity.Cultivation
	if err := r.db.Set("gorm:auto_preload", true).First(&cultivation, id).Error; err != nil {
		return nil, err
	}
	return &cultivation, nil
}

// FindAllByUID は指定されたUser.IDを持つ栽培している植物の一覧を取得します。
func (r *CultivationRepository) FindAllByUID(uid int) (*entity.Cultivations, error) {
	var cultivations entity.Cultivations
	if err := r.db.Set("gorm:auto_preload", true).Find(&cultivations, "uid = ?", uid).Error; err != nil {
		return nil, err
	}
	return &cultivations, nil
}

// Store は栽培している植物を新規保存します。
func (r *CultivationRepository) Store(cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	if err := r.db.Set("gorm:auto_preload", true).Create(&cultivation).Error; err != nil {
		return nil, err
	}
	return cultivation, nil
}

// UpdateByID は栽培している植物の情報を更新します。
func (r *CultivationRepository) UpdateByID(cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	if err := r.db.Set("gorm:auto_preload", true).Model(&entity.User{}).Update(&cultivation).First(&cultivation).Error; err != nil {
		return nil, err
	}
	return cultivation, nil
}

// DeleteByID は指定されたIDを持つ栽培している植物を削除します。
func (r *CultivationRepository) DeleteByID(id int) error {
	var cultivation entity.Cultivation
	if err := r.db.Where("id = ?", id).Delete(&cultivation).Error; err != nil {
		return err
	}
	return nil
}

// CheckByIDUID は栽培している植物にid, user_idの組のデータが存在すればtrue、なければfalseを返却します。
func (r *CultivationRepository) CheckByIDUID(id int, uid int) (bool, error) {
	var cultivations entity.Cultivations
	if err := r.db.Where("id = ? AND uid = ?", id, uid).Find(&cultivations).Error; err != nil {
		return false, err
	}
	return true, nil
}
