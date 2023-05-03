package server

import (
	"net/http"
)

type User struct {
	Id       int
	Nickname string
}

var (
	usrs []*User
)

func NewServer(addr string) *http.Server {
	InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
