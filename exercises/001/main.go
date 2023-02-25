package main

import "fmt"

type User struct {
	firstName string
	lastName  string
}

func main() {
	u := User{
		firstName: "Amirreza",
		lastName:  "Kheradmand",
	}

	fmt.Println(u.getName())
	fmt.Println(u.firstName)

}

func (u *User) getName() string {
	return u.firstName + " " + u.lastName
}
