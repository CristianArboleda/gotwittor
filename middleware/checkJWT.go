package middleware

import (
	"net/http"

	"github.com/CristianArboleda/gotwittor/routes"
)

// CheckJWT : check if the JWR is valid
func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "Error in Token: "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
