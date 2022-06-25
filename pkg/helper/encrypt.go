package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/adi221/good-reads/pkg/config"
	"github.com/adi221/good-reads/pkg/model"
)

func ValidateAndEncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password, err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(encryptedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func EncodeUser(user *model.User) (string, error) {
	jwtConfig := config.GetJWTConfig()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = jwtConfig.ExpiresTime().Unix()

	return token.SignedString([]byte(jwtConfig.Secret))
}

func ExtractClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return []byte(config.GetJWTConfig().Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}

func DecodeTokenFromRequest(req *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return []byte(config.GetJWTConfig().Secret), nil
	})

	return token, err
}
