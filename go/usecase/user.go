package usecase

import (
	"context"
	"fmt"

	"github.com/jphacks/F_2002_1/go/domain/entity"
	"github.com/jphacks/F_2002_1/go/domain/repository"
)

// UserUseCase はユーザに関係するアプリケーションロジックを担当する構造体です。
type UserUseCase struct {
	userRepo repository.User
}

// NewUserUseCase はUserUseCaseのポインタを生成する関数です。
func NewUserUseCase(userRepo repository.User) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// ReadUsers は全ユーザを取得します。
func (u *UserUseCase) ReadUsers(ctx context.Context) (*entity.Users, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("find users from repo: %w", err)
	}
	return users, nil
}

// ReadUser はユーザを取得します。
func (u *UserUseCase) ReadUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("find user from repo id=%s: %w", id, err)
	}
	return user, nil
}

// CreateUser はユーザを作成します。
func (u *UserUseCase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := u.userRepo.Store(user)
	if err != nil {
		return nil, fmt.Errorf("store user from repo: %w", err)
	}
	return user, nil
}

// UpdateUser はユーザを作成します。
func (u *UserUseCase) UpdateUser(ctx context.Context, id string, user *entity.User) (*entity.User, error) {
	user, err := u.userRepo.UpdateByID(id, user)
	if err != nil {
		return nil, fmt.Errorf("update user from repo id=%s: %w", id, err)
	}
	return user, nil
}

// DeleteUser はユーザを作成します。
func (u *UserUseCase) DeleteUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.userRepo.DeleteByID(id)
	if err != nil {
		return nil, fmt.Errorf("delete user from repo id=%s: %w", id, err)
	}
	return user, nil
}