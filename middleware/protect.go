package middleware

import (
	"net/http"
)

func Protect(w http.ResponseWriter, r *http.Request) {
	// authorization := r.Header.Get("Authorization")
	// separateBearer := strings.Replace(authorization, "Bearer", "", 1)
	// token = strings.TrimSpace(separateBearer)
	// return token
}
