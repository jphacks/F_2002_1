package entity

import "errors"

var (
	// ErrPlantNotFound は植物が存在しないエラーを表します。
	ErrPlantNotFound = errors.New("plant not found")
	// ErrUserNotFound はユーザが存在しないエラーを表します。
	ErrUserNotFound = errors.New("user not found")
	// ErrCultivationNotFound は栽培物が存在しないエラーを表します。
	ErrCultivationNotFound = errors.New("cultivation not found")

	ErrInvalidIdToken = errors.New("Invalid ID token")
)
