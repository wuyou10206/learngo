package basic

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := calTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calTriangle(%d,%d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
