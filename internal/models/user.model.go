package models

import (
	"fmt"
)

// User holds personal user information
type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserStorage defines all the database operations
type UserStorage interface {
	ListUsers() ([]User, error)
	GetUser(i int) (User, error)
	AddUser(u User) (User, error)
	UpdateUser(u User) (User, error)
	DeleteUser(i int) error
}

// GoString implements the GoStringer interface so we can display the full struct during debugging
// usage: fmt.Printf("%#v", i)
// ensure that i is a pointer, so might need to do &i in some cases
func (u *User) GoString() string {
	// extend to base goString

	return fmt.Sprintf(`
{	
	%s,
	FirstName: %s,
	LastName: %s,

}`,
		u.Base.GoString(),
		u.FirstName,
		u.LastName,
	)
}
