package hello_world

import "testing"

func TestSayHello(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"With name", "Alice", "Hello, Alice!"},
		{"Empty name", "", "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SayHello(tt.input)
			if got != tt.expected {
				t.Errorf("SayHello(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
