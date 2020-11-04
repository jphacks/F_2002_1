package usecase

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/database"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"
)

// CultivationUseCase は栽培物に関係するアプリケーションロジックを担当する構造体です。
type CultivationUseCase struct {
	cultivationRepo repository.Cultivation
}

// NewCultivationUseCase はCultivationUseCaseのポインタを生成する関数です。
func NewCultivationUseCase(db *gorm.DB) *CultivationUseCase {
	return &CultivationUseCase{cultivationRepo: database.NewCultivationRepository(db)}
}

// ReadCultivation は栽培物を取得します。
func (c *CultivationUseCase) ReadCultivation(id int) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("find cultivation from repo id=%s: %w", id, err)
	}
	return cultivation, nil
}

// ReadCultivationsByUid はあるユーザの全栽培物を取得します。
func (c *CultivationUseCase) ReadCultivationByUUid(uuid string, id int) (*entity.Cultivations, error) {
	cultivations, err := c.cultivationRepo.FindByUUidId(uuid, id)
	if err != nil {
		return nil, fmt.Errorf("find cultivations from repo: %w", err)
	}
	return cultivations, nil
}

// CreateCultivation はあるユーザの栽培物を作成します。
func (c *CultivationUseCase) CreateCultivationByUid(uid int, cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.Store(uid, cultivation)
	if err != nil {
		return nil, fmt.Errorf("store cultivation from repo: %w", err)
	}
	return cultivation, nil
}

// CreateCultivation はc.を作成します。
func (c *CultivationUseCase) CreateCultivationByUuid(uuid string, cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.StoreByUuid(uuid, cultivation)
	if err != nil {
		return nil, fmt.Errorf("store cultivation from repo: %w", err)
	}
	return cultivation, nil
}

// UpdateCultivation は栽培物を作成します。
func (c *CultivationUseCase) UpdateCultivation(id int, cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.UpdateByID(id, cultivation)
	if err != nil {
		return nil, fmt.Errorf("update cultivation from repo id=%s: %w", id, err)
	}
	return cultivation, nil
}

// UpdateCultivation は栽培物を作成します。
func (c *CultivationUseCase) UpdateCultivationByUuid(uuid string, id int, cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.UpdateByID(uuid, id, cultivation)
	if err != nil {
		return nil, fmt.Errorf("update cultivation from repo id=%s: %w", id, err)
	}
	return cultivation, nil
}

// DeleteCultivation は栽培物を作成します。
func (c *CultivationUseCase) DeleteCultivation(id int) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.DeleteByID(id)
	if err != nil {
		return nil, fmt.Errorf("delete cultivation from repo id=%s: %w", id, err)
	}
	return cultivation, nil
}


// DeleteCultivation は栽培物を作成します。
func (c *CultivationUseCase) DeleteCultivationByUuid(uuid string, id int) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.DeleteByUidId(uuid, id)
	if err != nil {
		return nil, fmt.Errorf("delete cultivation from repo id=%s: %w", id, err)
	}
	return cultivation, nil
}
