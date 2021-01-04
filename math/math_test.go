package math

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(3, 4)
	if sum != 7 {
		t.Errorf("Wanted 7 but received %d", sum)
	}
}
