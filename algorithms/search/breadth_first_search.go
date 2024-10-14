package search

import "grokking-algorithms/algorithms/utils"

func BreadthFirstSearch(graph map[string][]string, start string, target string) bool {
	searchQueue := graph[start]
	var searched []string

	for len(searchQueue) > 0 {
		person := searchQueue[0]
		searchQueue = searchQueue[1:]

		if !utils.Contains(searched, person) {
			if person == target {
				return true
			} else {
				searchQueue = append(searchQueue, graph[person]...)
				searched = append(searched, person)
			}
		}
	}

	return false
}
