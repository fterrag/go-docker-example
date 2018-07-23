package user

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetFullName(t *testing.T) {
	u := User{NameFirst: "John", NameLast: "Smith"}

	expected := "John Smith"
	actual := u.GetFullName()

	assert.Equal(t, expected, actual)
}
