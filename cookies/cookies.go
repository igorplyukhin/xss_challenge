package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetCookie(cookieName string, request *http.Request) (val string) {
	if cookie, err := request.Cookie(cookieName); err == nil {
		cookieValue := ""
		if err = cookieHandler.Decode(cookieName, cookie.Value, &cookieValue); err == nil {
			return cookieValue
		}
	}
	return val
}

func SetCookie(cookieName string, val string, response http.ResponseWriter) {
	value := val
	if encoded, err := cookieHandler.Encode(cookieName, value); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
			Expires: time.Now().Add(time.Hour * 24 * 30),
		}
		http.SetCookie(response, cookie)
	}
}

func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
