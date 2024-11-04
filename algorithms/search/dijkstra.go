package search

import "math"

// Dijkstra finds the shortest path from the start node to all other nodes in the graph.
func Dijkstra(graph map[string]map[string]int, start string) (map[string]int, map[string]string) {
	costs := make(map[string]int)
	parents := make(map[string]string)
	processed := make(map[string]bool)

	// Initialize costs and parents
	for node := range graph {
		if node == start {
			costs[node] = 0
		} else {
			costs[node] = math.MaxInt64 // Use a large number to represent infinity
		}
		parents[node] = ""
	}

	// Process each node
	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n, c := range neighbors {
			newCost := cost + c
			if newCost < costs[n] {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode(costs, processed)
	}

	return costs, parents
}

// findLowestCostNode finds the node with the lowest cost that hasn't been processed yet.
func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt64
	var lowestCostNode string
	for node, cost := range costs {
		if cost < lowestCost && !processed[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}
