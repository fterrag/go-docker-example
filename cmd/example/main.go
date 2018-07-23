package main

import (
	"log"

	"github.com/fterrag/go-docker-example/user"

	"github.com/fatih/color"
)

func main() {
	u := user.User{NameFirst: "John", NameLast: "Smith"}

	greeting := color.GreenString("Hello %s!\n", u.GetFullName())
	log.Printf(greeting)
}
