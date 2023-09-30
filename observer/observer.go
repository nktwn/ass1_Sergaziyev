package observer

import "fmt"

type Observer interface {
	Update(name string, code string, teacher string)
	GetEmail() string
}

type User struct {
	email    string
	subjects []string
}

func NewUser(email string, subjects []string) *User {
	return &User{
		email:    email,
		subjects: subjects,
	}
}

func (u *User) Update(name string, code string, teacher string) {
	fmt.Println("New grade for subject:", name, code, teacher)
}

func (u *User) GetEmail() string {
	return u.email
}
