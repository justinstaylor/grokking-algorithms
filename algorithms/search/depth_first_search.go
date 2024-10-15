package search

func DepthFirstSearch(graph map[string][]string, start string, target string) bool {
	if start == target {
		return true
	}

	for _, v := range graph[start] {
		if DepthFirstSearch(graph, v, target) {
			return true
		}
	}

	return false
}
