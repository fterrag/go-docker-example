package main

import (
	"log"

	"github.com/fterrag/go-docker-example/user"
)

func main() {
	u := user.User{NameFirst: "John", NameLast: "Smith"}

	log.Printf("Hello %s!\n", u.GetFullName())
}
