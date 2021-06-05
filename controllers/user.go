package controllers

import (
	"log"
	"net/http"
)

type User struct {
	l *log.Logger
}

func (u *User) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rw.Write([]byte("Yo there"))
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func NewUser(l *log.Logger) *User {
	return &User{l}
}
