package hello_world

// SayHello returns a greeting message.
func SayHello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
