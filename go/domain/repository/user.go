package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// User はユーザの永続化を担当するリポジトリです。
type User interface {
	FindByID(id int) (*entity.User, error)
	FindIDByUID(uid string) (int, error)
	FindAll() (*entity.Users, error)
	Store(user *entity.User) (*entity.User, error)
	UpdateByID(user *entity.User) (*entity.User, error)
	DeleteByID(id int) error
}
