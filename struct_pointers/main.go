package main

import (
	"fmt"

	"github.com/your-username/go-struct-pointer-example/examples/struct_pointers"
)

func main() {
	// Create a new user
	user := struct_pointers.NewUser("Charlie", 40)

	// Print greeting
	fmt.Println(user.Greet())

	// Celebrate birthday
	user.Birthday()
	fmt.Printf("%s is now %d years old.\n", *user.Name, *user.Age)
}
