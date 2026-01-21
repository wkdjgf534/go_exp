package math_test

import (
	"os"
	"testing"

	math "lesson-66/math"
)

var something string

// This function execites before other tests
func TestMain(m *testing.M) {
	something = "test"
	exitCode := m.Run()

	os.Exit(exitCode)
}

func Test_MinInt(t *testing.T) {
	t.Cleanup(func() {
		// Any cleanup here
	})
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "100 is less than 1000",
			a:        100,
			b:        1000,
			expected: 100,
		},
		{
			name:     "-2 is less than -1",
			a:        -1,
			b:        -2,
			expected: -2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := math.MinInt(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("expected: %d, got: %d", tc.expected, result)
			}
		})
	}
}

func Test_MaxInt(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "10 is more than 5",
			a:        5,
			b:        10,
			expected: 10,
		},
		{
			name:     "-5 is more than -10",
			a:        -5,
			b:        -10,
			expected: -5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := math.MaxInt(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected: %d, got: %d", tc.expected, result)
			}
		})
	}
}
