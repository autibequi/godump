package string

import "testing"

func Test(t *testing.T) {
	// Inicia fixtures
	var tests = []struct {
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"😋i", "i😋"},
	}

	// Implementa de fato os testes
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
