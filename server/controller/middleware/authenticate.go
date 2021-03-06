package middleware

import (
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/standup-raven/standup-raven/server/config"
)

// Authenticate middleware verifies the request was made by a logged in Mattermost user.
// this is checked by the presence of Mattermost-User-Id HTTP header.
func Authenticate(w http.ResponseWriter, r *http.Request) *model.AppError {
	userID := r.Header.Get(config.HeaderMattermostUserID)

	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return model.NewAppError("MiddlewareAuthenticate", "", nil, "Unauthorized", http.StatusUnauthorized)
	}
	return nil
}
