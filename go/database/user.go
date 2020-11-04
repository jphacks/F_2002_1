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
func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	return nil, nil
}

// FindAll は指定されたIDを持つユーザを取得します。
func (r *UserRepository) FindAll() (*entity.Users, error) {
	return nil, nil
}
// Store はユーザを新規保存します。
func (r *UserRepository) Store(user *entity.User) (*entity.User, error) {
	return nil, nil
}

// Update はユーザの情報を更新します。
func (r *UserRepository) UpdateByID(id string, user *entity.User) (*entity.User, error) {
	return nil, nil
}

// DeleteByID は指定されたIDを持つユーザを削除します。
func (r *UserRepository) DeleteByID(id string) (*entity.User, error) {
	return nil, nil
}