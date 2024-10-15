package utils

import "testing"

func TestContains(t *testing.T) {
	arr := []string{"a", "b", "c"}
	item := "b"
	want := true
	got := Contains(arr, item)

	if got != want {
		t.Errorf("Contains(%v, %v) = %v; want %v", arr, item, got, want)
	}
}
