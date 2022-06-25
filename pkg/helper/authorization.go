package helper

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte("very-secret")
var s = securecookie.New(hashKey, nil)
var cookieName = "cookie"

func SetCookieHandler(w http.ResponseWriter, cookieValue string) (string, error) {
	encoded, err := s.Encode(cookieName, cookieValue)
	if err != nil {
		return "", err
	}
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   encoded,
		Expires: time.Now().Add(time.Hour * 24),
	}
	http.SetCookie(w, cookie)
	return encoded, err
}

func ReadCookieHandler(r *http.Request) (string, error) {
	var value string
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	if err = s.Decode(cookieName, cookie.Value, &value); err != nil {
		return "", err
	}
	return value, nil
}

func AddTokenToResponseHeader(w http.ResponseWriter, tokenString string) {
	w.Header().Set("authorization", tokenString)
}
