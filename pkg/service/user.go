package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/adi221/good-reads/pkg/constant"
	"github.com/adi221/good-reads/pkg/helper"
	"github.com/adi221/good-reads/pkg/model"
)

func addCookieToRequest(ctx context.Context, tokenString string) error {
	type funcHandlerType func(str string)
	setCookieHandler := ctx.Value(constant.ContextCookieSetter)
	if f, ok := setCookieHandler.(func(str string)); ok {
		funcHandlerType(f)(tokenString)
	} else {
		return errors.New("Invalid cookie handler")
	}
	return nil
}

func addTokenToResponseHeader(ctx context.Context, tokenString string) error {
	type funcHandlerType func(str string)
	setTokenHandler := ctx.Value(constant.ContextAuthorizationTokenSetter)
	if f, ok := setTokenHandler.(func(str string)); ok {
		funcHandlerType(f)(tokenString)
	} else {
		return errors.New("Invalid authorization handler")
	}
	return nil
}

func (reg *Registry) SignUpUser(ctx context.Context, user model.User) (*model.User, error) {
	reg.logger.Debug().Msg("Creating user")
	newUser, err := reg.db.CreateUser(user)
	if err != nil {
		return nil, err
	}
	tokenString, err := helper.EncodeUser(newUser)
	addTokenToResponseHeader(ctx, tokenString)
	return newUser, nil
}

func (reg *Registry) LoginUser(ctx context.Context, identity string, password string) (*model.User, error) {
	reg.logger.Debug().Msg("Trying to log in a user")
	user, err := reg.db.GetUserByIdentityAndVerify(identity, password)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, helper.NewHttpError(
			model.UserNonExist,
			http.StatusNotFound,
			"User does not exist",
		)
	}
	tokenString, err := helper.EncodeUser(user)
	addTokenToResponseHeader(ctx, tokenString)
	return user, nil
}
