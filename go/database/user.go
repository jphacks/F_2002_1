package database

import (
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"

	"github.com/jinzhu/gorm"
)

var _ repository.User = &UserRepository{}

// UserRepository は repository.UserRepository を満たす構造体です。
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository はUserRepositoryのポインタを生成する関数です。
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByID は指定されたIDを持つユーザを取得します。
func (r *UserRepository) FindByID(id int) (*entity.User, error) {
	var user entity.User
	if err := r.db.Preload("Cultivations").Preload("Cultivations.Plant").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindIDByUID は指定されたUIDを持つユーザIDを取得します。
func (r *UserRepository) FindIDByUID(uid string) (int, error) {
	var user entity.User
	if err := r.db.Preload("Cultivations").Preload("Cultivations.Plant").Find(&user, "uid = ?", uid).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// FindAll は指定されたIDを持つユーザを取得します。
func (r *UserRepository) FindAll() (*entity.Users, error) {
	var users entity.Users
	if err := r.db.Preload("Cultivations").Preload("Cultivations.Plant").Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

// Store はユーザを新規保存します。
func (r *UserRepository) Store(user *entity.User) (*entity.User, error) {
	if err := r.db.Preload("Cultivations").Preload("Cultivations.Plant").Create(&user).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateByID はユーザの情報を更新します。
func (r *UserRepository) UpdateByID(user *entity.User) (*entity.User, error) {
	if err := r.db.Preload("Cultivations").Preload("Cultivations.Plant").Model(&entity.User{}).Update(&user).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteByID は指定されたIDを持つユーザを削除します。
func (r *UserRepository) DeleteByID(id int) error {
	if err := r.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
