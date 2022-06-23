package service

import (
	"github.com/adi221/good-reads/pkg/model"
)

func (reg *Registry) SignUpUser(user model.User) (*model.User, error) {
	reg.logger.Debug().Msg("Creating user")
	newUser, err := reg.db.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
