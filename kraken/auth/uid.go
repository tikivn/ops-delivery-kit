package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeExternalUIDMiddleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid := uidFromKrakenRequest(r)

			ctx := ContextWithUID(r.Context(), uid)

			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		})
	}
}

func uidFromKrakenRequest(r *http.Request) UID {
	userID := r.Header.Get(HeaderXuid)
	userTrn := r.Header.Get(HeaderXtrn)
	userEmail := r.Header.Get(HeaderXEmail)
	username := r.Header.Get(HeaderXUserName)

	return UID{ID: userID,
		Trn:      userTrn,
		Email:    userEmail,
		Username: username,
	}
}
