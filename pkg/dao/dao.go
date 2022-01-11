package dao

import "hackaton/pkg/model"

type DB interface {
	// User module
	CreateUser(user model.User) error
	GetUsers() (users []model.User, err error)

	Health() error
}
