package middleware

import (
	"net/http"

	"github.com/CristianArboleda/gotwittor/db"
)

/*CheckDB: check if the DB is active and redirect the request */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(rw, "DB connection Error", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
