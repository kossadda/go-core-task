package data

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		wantErr  bool
	}{
		{
			name:     "Valid input",
			input:    "apple banana cherry\n",
			expected: []string{"apple", "banana", "cherry"},
			wantErr:  false,
		},
		{
			name:     "Empty input",
			input:    "\n",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Whitespace input",
			input:    "    \n",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Single word",
			input:    "apple\n",
			expected: []string{"apple"},
			wantErr:  false,
		},
		{
			name:     "Special characters",
			input:    "apple @#$% ^&*()\n",
			expected: []string{"apple", "@#$%", "^&*()"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe: %v", err)
			}
			defer r.Close()

			_, _ = w.WriteString(tt.input)
			_ = w.Close()

			oldStdin := os.Stdin
			os.Stdin = r
			defer func() { os.Stdin = oldStdin }()

			result, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error: %v, got: %v", tt.wantErr, err)
			}

			if !equal(result, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
