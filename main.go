package main

import (
	"fmt"

	"github.com/fterrag/go-docker-example/user"
)

func main() {
	u := user.User{NameFirst: "John", NameLast: "Smith"}

	fmt.Printf("Hello %s!\n", u.GetFullName())
}
