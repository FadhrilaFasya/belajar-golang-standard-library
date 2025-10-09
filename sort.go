package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User // Define a new type UserSlice which is a slice of User

func (u UserSlice) Len() int { //
	return len(u)
}

func (u UserSlice) Less(i, j int) bool { // Compare the elements with indexes i and j
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) { // Swap the elements with indexes i and j
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"John", 30},
		{"Jane", 25},
		{"Doe", 35},
		{"Smith", 28},
	}

	sort.Sort(UserSlice(users)) // Sort the users slice using the UserSlice type

	fmt.Println("Sorted by Age:")
	for _, user := range users { // Iterate over the sorted users slice
		fmt.Printf("%s is %d years old\n", user.Name, user.Age)
	}
}