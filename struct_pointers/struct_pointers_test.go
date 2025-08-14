package struct_pointers

import "testing"

func TestNewUserAndGreet(t *testing.T) {
	user := NewUser("Alice", 25)

	got := user.Greet()
	want := "Hello, Alice!"

	if got != want {
		t.Errorf("Greet() = %q; want %q", got, want)
	}
}

func TestBirthday(t *testing.T) {
	user := NewUser("Bob", 30)
	user.Birthday()

	if *user.Age != 31 {
		t.Errorf("Expected age 31, got %d", *user.Age)
	}
}
