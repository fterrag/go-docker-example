package user

import "fmt"

// A User represents a user.
type User struct {
	NameFirst string
	NameLast  string
}

// GetFullName returns the first and last name of the user.
func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.NameFirst, u.NameLast)
}
