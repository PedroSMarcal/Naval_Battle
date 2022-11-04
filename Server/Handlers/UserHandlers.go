package handlers

import (
	"io"
	"net/http"
)

type userHandlers struct{}

func NewUserHandler() userHandlers {
	return userHandlers{}
}

func (u *userHandlers) UserGet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "getUser")
}
