package config

import (
	"time"
)

var oneMonthInMinutes int64 = 43200

func GetJWTConfig() JWTConfig {
	return JWTConfig{
		Secret:       "good-reads secert key",
		ValidMinutes: oneMonthInMinutes,
		Realm:        "jwt auth",
	}
}

func (jwt *JWTConfig) ExpiresTime() time.Time {
	minutes := time.Minute * time.Duration(jwt.ValidMinutes)
	return time.Now().Add(minutes)
}
