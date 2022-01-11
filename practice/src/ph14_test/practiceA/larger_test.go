package practiceA

import (
	"fmt"
	"testing"
)

func TestFirstLarge(t *testing.T) {
	want := 2
	got := Larger(2, 1)
	if want != got {
		t.Error(errorString(2, 1, got, want))
	}
}

func TestSecondLarge(t *testing.T) {
	want := 8
	got := Larger(4, 8)
	if want != got {
		t.Error(errorString(4, 8, got, want))
	}
}

func errorString(a int, b int, got int, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = %d, want %d", a, b, got, want)
}
