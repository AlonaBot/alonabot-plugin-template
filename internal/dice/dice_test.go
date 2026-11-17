package dice

import "testing"

func TestRoll(t *testing.T) {
	for range 1000 {
		value, ok := Roll(6)
		if !ok || value < 1 || value > 6 {
			t.Fatalf("Roll(6) = %d, %v", value, ok)
		}
	}
	for _, sides := range []int64{1, 101} {
		if _, ok := Roll(sides); ok {
			t.Errorf("Roll(%d) ok = true, want false", sides)
		}
	}
}
