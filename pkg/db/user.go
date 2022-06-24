package db

import (
	"github.com/adi221/good-reads/pkg/model"
)

type UserRespository interface {
	CreateUser(user model.User) (*model.User, error)
	GetUserByIdentity(identity string, password string) (*model.User, error)
}
