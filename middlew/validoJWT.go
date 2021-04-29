package middlew

import (
	"net/http"

	"github.com/francovico/twittor/routers"
)

func ValidoJWT(http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token "+err.Error(), http.StatusBadRequest)
			return
		}
	}
}
