package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetCookie(cookieName string, key string,  request *http.Request) (val string) {
	if cookie, err := request.Cookie(cookieName); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(cookieName, cookie.Value, &cookieValue); err == nil {
			a := cookieValue
			val = a[key]
		}
	}
	return val
}

func SetSession(cookieName string, key string,val string, response http.ResponseWriter) {
	value := map[string]string{
		key: val,
	}
	if encoded, err := cookieHandler.Encode(cookieName, value); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
