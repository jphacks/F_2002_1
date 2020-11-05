package usecase

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/database"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"
)

// PlantUseCase は植物に関係するアプリケーションロジックを担当する構造体です。
type PlantUseCase struct {
	plantRepo repository.Plant
}

// NewPlantUseCase はPlantUseCaseのポインタを生成する関数です。
func NewPlantUseCase(db *gorm.DB) *PlantUseCase {
	return &PlantUseCase{plantRepo: database.NewPlantRepository(db)}
}

// ReadPlants は全植物を取得します。
func (p *PlantUseCase) ReadPlants() (*entity.Plants, error) {
	plants, err := p.plantRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("find plants from repo: %w", err)
	}
	return plants, nil
}

// // ReadPlant は植物を取得します。
// func (p *PlantUseCase) ReadPlant(id int) (*entity.Plant, error) {
// 	plant, err := p.plantRepo.FindByID(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("find plant from repo id=%s: %w", id, err)
// 	}
// 	return plant, nil
// }

// // CreatePlant は植物を作成します。
// func (p *PlantUseCase) CreatePlant(plant *entity.Plant) (*entity.Plant, error) {
// 	plant, err := p.plantRepo.Store(plant)
// 	if err != nil {
// 		return nil, fmt.Errorf("store plant from repo: %w", err)
// 	}
// 	return plant, nil
// }

// // UpdatePlant は植物を編集します。
// func (p *PlantUseCase) UpdatePlant(plant *entity.Plant) (*entity.Plant, error) {
// 	plant, err := p.plantRepo.UpdateByID(id, plant)
// 	if err != nil {
// 		return nil, fmt.Errorf("update plant from repo id=%s: %w", plant.ID, err)
// 	}
// 	return plant, nil
// }

// // DeletePlant は植物を削除します。
// func (p *PlantUseCase) DeletePlant(id int) (*entity.Plant, error) {
// 	plant, err := p.plantRepo.DeleteByID(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("delete plant from repo id=%s: %w", id, err)
// 	}
// 	return plant, nil
// }
