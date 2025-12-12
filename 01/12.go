package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateAge(u *User, newAge int) {
	u.Age = newAge
}

func main() {
	u := User{"Sandeep", 23}
	updateAge(&u, 25)
	fmt.Println(u)
}
