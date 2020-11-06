package usecase

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jphacks/F_2002_1/go/database"
	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"
)

// UserUseCase はユーザに関係するアプリケーションロジックを担当する構造体です。
type UserUseCase struct {
	userRepo repository.User
}

// NewUserUseCase はUserUseCaseのポインタを生成する関数です。
func NewUserUseCase(db *gorm.DB) *UserUseCase {
	return &UserUseCase{userRepo: database.NewUserRepository(db)}
}

// ReadUsers は全ユーザを取得します。
func (u *UserUseCase) ReadUsers() (*entity.Users, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("find users from repo: %w", err)
	}
	return users, nil
}

// ReadUser はユーザを取得します。
func (u *UserUseCase) ReadUser(id int) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("find user from repo id=%d: %w", id, err)
	}
	return user, nil
}

// ReadUserIDByUID はユーザを取得します。
func (u *UserUseCase) ReadUserIDByUID(uid string) (int, error) {
	id, err := u.userRepo.FindIDByUID(uid)
	if err != nil {
		return 0, fmt.Errorf("find user from repo uid=%s: %w", uid, err)
	}
	return id, nil
}

// CreateUser はユーザを作成します。
func (u *UserUseCase) CreateUser(user *entity.User) (*entity.User, error) {
	user, err := u.userRepo.Store(user)
	if err != nil {
		return nil, fmt.Errorf("store user from repo: %w", err)
	}
	return user, nil
}

// UpdateUser はユーザを作成します。
func (u *UserUseCase) UpdateUser(user *entity.User) (*entity.User, error) {
	user, err := u.userRepo.UpdateByID(user)
	if err != nil {
		return nil, fmt.Errorf("update user from repo id=%d: %w", user.ID, err)
	}
	return user, nil
}

// DeleteUser はユーザを作成します。
func (u *UserUseCase) DeleteUser(id int) error {
	err := u.userRepo.DeleteByID(id)
	if err != nil {
		return fmt.Errorf("delete user from repo id=%d: %w", id, err)
	}
	return nil
}
