package user

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetFullName(t *testing.T) {
	assert := assert.New(t)

	u := User{NameFirst: "John", NameLast: "Smith"}

	expected := "John Smith"
	actual := u.GetFullName()

	assert.Equal(expected, actual)
}
