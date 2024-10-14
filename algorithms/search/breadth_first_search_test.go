package search

import "testing"

func TestBreadthFirstSearch(t *testing.T) {
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
	got := BreadthFirstSearch(graph, start, target)

	if got != want {
		t.Errorf("BreadthFirstSearch(%v, %v, %v) = %v; want %v", graph, start, target, got, want)
	}
}
