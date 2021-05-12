package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

// From docs:
// GenerateRandomKey creates a random key with the given length in bytes.
// On failure, returns nil.
//
// Note that keys created using `GenerateRandomKey()` are not automatically
// persisted. New keys will be created when the application is restarted, and
// previously issued cookies will not be able to be decoded.
//
// Callers should explicitly check for the possibility of a nil return, treat
// it as a failure of the system random number generator, and not continue.
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetCookie(cookieName string, request *http.Request) (val string) {
	if cookie, err := request.Cookie(cookieName); err == nil {
		cookieValue := ""
		err = cookieHandler.Decode(cookieName, cookie.Value, &cookieValue);
		if err == nil {
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
