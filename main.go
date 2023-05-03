package main

import (
	"github.com/fixztter/rest-fundamentals/server"
)

func main() {
	sv := server.NewServer(":3000")
	err := sv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
