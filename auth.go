package main

import (
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func (auth *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.next.ServeHTTP(w, r)
}

func MustAuth(h http.Handler) http.Handler {
	return &authHandler{next: h}
}
