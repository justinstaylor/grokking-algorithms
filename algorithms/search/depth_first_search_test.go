package search

import "testing"

func TestDepthFirstSearch(t *testing.T) {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	start := "you"
	target := "jonny"
	want := true
	got := DepthFirstSearch(graph, start, target)

	if got != want {
		t.Errorf("DepthFirstSearch(%v, %v, %v) = %v; want %v", graph, start, target, got, want)
	}
}
