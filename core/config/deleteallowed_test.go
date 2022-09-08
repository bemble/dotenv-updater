package config

import (
	"os"
	"testing"
)

func TestIsDeleteAllowed(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"1", true},
		{"True", true},
		{"true", true},
		{"0", false},
		{"tRue", false},
		{"tRue", false},
	}

	for _, tc := range cases {
		t.Run(tc.value, func(t *testing.T) {
			os.Setenv("DELETE_ALLOWED", tc.value)

			isNoDelete := IsDeleteAllowed()

			if tc.expected != isNoDelete {
				t.Fatalf("Expected is no delete %t, but got %t", tc.expected, isNoDelete)
			}
		})
	}
}
