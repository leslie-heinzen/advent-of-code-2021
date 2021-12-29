package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	name      string
	neighbors map[string]bool
}

func addUpdateNode(a string, b string, graph map[string]node) {
	if _, ok := graph[a]; !ok {
		graph[a] = node{
			name:      a,
			neighbors: map[string]bool{b: true},
		}
	} else {
		if _, ok := graph[a].neighbors[b]; !ok {
			graph[a].neighbors[b] = true
		}
	}
}

func findAllPaths(graph map[string]node, part2 bool) int {
	paths := [][]string{}

	var mapRoute func(name string, pathProgress []node)
	mapRoute = func(name string, pathProgress []node) {
		currentNode := graph[name]
		pathProgress = append(pathProgress, currentNode)

		visited := map[string]int{}

		for _, p := range pathProgress {
			visited[p.name]++
		}

		if currentNode.name == "end" {
			path := []string{}

			for _, p := range pathProgress {
				path = append(path, p.name)
			}

			paths = append(paths, path)
			return
		}

		for k := range currentNode.neighbors {
			next := graph[k]
			if next.name == "start" {
				continue
			}

			isSmallCave := func(a string) bool { return strings.ToLower(a) == a }

			if isSmallCave(next.name) && visited[next.name] >= 1 {
				if !part2 {
					continue
				} else {
					skip := false
					for k := range visited {
						if isSmallCave(graph[k].name) && visited[k] >= 2 {
							skip = true
							break
						}
					}

					if skip {
						continue
					}
				}
			}

			mapRoute(k, pathProgress)
		}
	}

	mapRoute("start", []node{})
	return len(paths)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	graph := map[string]node{}
	for _, line := range strings.Fields(string(input)) {
		parts := strings.Split(line, "-")
		addUpdateNode(parts[0], parts[1], graph)
		addUpdateNode(parts[1], parts[0], graph)
	}

	solution1 := findAllPaths(graph, false)
	solution2 := findAllPaths(graph, true)

	fmt.Printf("Solution 1: %d\n", solution1)
	fmt.Printf("Solution 2: %d\n", solution2)
}
