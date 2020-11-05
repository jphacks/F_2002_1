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

// ReadCultivationsByUID はあるユーザの全ての栽培物を取得します。
func (c *CultivationUseCase) ReadCultivationsByUID(uid int) (*entity.Cultivations, error) {
	cultivations, err := c.cultivationRepo.FindAllByUID(uid)
	if err != nil {
		return nil, fmt.Errorf("find cultivation from repo user_id=%s: %w", uid, err)
	}
	return cultivations, nil
}

// ReadCultivation は栽培物を取得します。
func (c *CultivationUseCase) ReadCultivation(id int) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("find cultivation from repo id=%s: %w", id, err)
	}
	return cultivation, nil
}

// ReadCultivation は栽培物を取得します。
func (c *CultivationUseCase) CheckCultivationByIDUID(id int, uid int) (bool, error) {
	isExist, err := c.cultivationRepo.CheckByIDUID(id, uid)
	if err != nil {
		return false, fmt.Errorf("find cultivation from repo user_id=%s: %w", uid, err)
	}
	return isExist, nil
}

// CreateCultivation はあるユーザの栽培物を作成します。
func (c *CultivationUseCase) CreateCultivation(cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.Store(cultivation)
	if err != nil {
		return nil, fmt.Errorf("store cultivation from repo: %w", err)
	}
	return cultivation, nil
}

// UpdateCultivation は栽培物を作成します。
func (c *CultivationUseCase) UpdateCultivation(cultivation *entity.Cultivation) (*entity.Cultivation, error) {
	cultivation, err := c.cultivationRepo.UpdateByID(cultivation)
	if err != nil {
		return nil, fmt.Errorf("update cultivation from repo id=%s: %w", cultivation.ID, err)
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
