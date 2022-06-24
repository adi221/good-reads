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

func (reg *Registry) LoginUser(identity string, password string) (*model.User, error) {
	reg.logger.Debug().Msg("Trying to log in a user")
	user, err := reg.db.GetUserByIdentityAndVerify(identity, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
