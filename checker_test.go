package disposableemail

import "testing"

func TestIsDisposable(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@10minutemail.com", true},
		{"user@gmail.com", false},
		{"invalid-email", false},
	}

	for _, tt := range tests {
		result := IsDisposable(tt.email)
		if result != tt.expected {
			t.Errorf("IsDisposable(%s) = %v, want %v", tt.email, result, tt.expected)
		}
	}
}