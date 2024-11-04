package search

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := map[string]map[string]int{
		"A": {"B": 1, "C": 4},
		"B": {"C": 2, "D": 5},
		"C": {"D": 1},
		"D": {},
	}

	expectedCosts := map[string]int{
		"A": 0,
		"B": 1,
		"C": 3,
		"D": 4,
	}

	expectedParents := map[string]string{
		"B": "A",
		"C": "B",
		"D": "C",
	}

	costs, parents := Dijkstra(graph, "A")

	for node, cost := range expectedCosts {
		if costs[node] != cost {
			t.Errorf("Expected cost for node %s: %d, got: %d", node, cost, costs[node])
		}
	}

	for node, parent := range expectedParents {
		if parents[node] != parent {
			t.Errorf("Expected parent for node %s: %s, got: %s", node, parent, parents[node])
		}
	}
}
