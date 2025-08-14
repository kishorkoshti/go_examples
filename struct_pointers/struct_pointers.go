package struct_pointers

import "fmt"

// User represents a person with a name and age stored as a pointer.
type User struct {
	Name *string
	Age  *int
}

// NewUser creates a new User instance using pointers.
func NewUser(name string, age int) User {
	return User{
		Name: &name,
		Age:  &age,
	}
}

// Birthday increases the user's age by 1 (pointer mutation).
func (u *User) Birthday() {
	if u.Age != nil {
		*u.Age++
	}
}

// Greet returns a greeting message.
func (u User) Greet() string {
	if u.Name == nil {
		return "Hello, Guest!"
	}
	return fmt.Sprintf("Hello, %s!", *u.Name)
}
