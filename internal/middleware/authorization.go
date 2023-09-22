package middleware

import (
	"errors"
	"net/http"

	"github.com/tylerb890/goapi/api"
	"github.com/tylerb890/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var NotAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(NotAuthorizedError)
			api.RequestErrorHandler(w, NotAuthorizedError)
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

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(NotAuthorizedError)
			api.RequestErrorHandler(w, NotAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}