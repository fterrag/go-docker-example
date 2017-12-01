package user

import "testing"

func TestGetFullName(t *testing.T) {
	u := User{NameFirst: "John", NameLast: "Smith"}

	expected := "John Smith"
	actual := u.GetFullName()

	if expected != actual {
		t.Errorf("GetFullName was incorrect, expected: %s, actual: %s.", expected, actual)
	}
}
