package middleware

import (
	"errors"
	"net/http"

	"goapi/api"
	"goapi/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedErrorEmpty = errors.New("Empty username or token.")
var UnAuthorizedErrorAuth = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedErrorEmpty)
			api.RequestErrorHandler(w, UnAuthorizedErrorEmpty)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedErrorAuth)
			api.RequestErrorHandler(w, UnAuthorizedErrorAuth)
			return
		}

		next.ServeHTTP(w, r)
	})
}
