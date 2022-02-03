package flash

import (
	"net/http"
	"time"
)

func SetFlash(w http.ResponseWriter, name string, value string) {
	c := &http.Cookie{Name: name, Value: value}
	http.SetCookie(w, c)
}

func GetFlash (w http.ResponseWriter, r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return "", nil
		default:
			return "", err
		}
	}
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 2)}
	http.SetCookie(w, dc)
	return c.Value, nil
}