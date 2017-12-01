package user

// A User represents a user.
type User struct {
	NameFirst string
	NameLast  string
}

// GetFullName returns the first name and last name of the user.
func (u *User) GetFullName() string {
	return u.NameFirst + " " + u.NameLast
}
