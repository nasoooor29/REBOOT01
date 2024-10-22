package middleware

import (
	"db-test/handlers"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			errMsg, good := err.(string)
			if !good {
				handlers.HandleErrorPage(http.StatusInternalServerError, "Error 500", "", "Internal server error")
				return
			}
			handlers.HandleErrorPage(http.StatusInternalServerError, "Error 500", "", errMsg)
		}()

		next.ServeHTTP(w, r)
	})
}
