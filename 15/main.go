package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	x                int
	y                int
	riskLevel        int
	terminus         bool
	terminusDistance int
	score            int
	neighbors        []*node
}

type xy struct {
	x int
	y int
}

// DIY priority queue interface
type queue []*node

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].score < q[j].score
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x interface{}) {
	node := x.(*node)
	*q = append(*q, node)
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*q = old[:n-1]
	return node
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fields := strings.Fields(string(input))
	nodes := defineNodes(fields)
	assignNeighbors(nodes)
	path := findPath(nodes)

	totalRisk := 0

	for _, node := range path {
		totalRisk += node.riskLevel
	}

	fmt.Print(totalRisk)
}

func defineNodes(fields []string) map[xy]*node {
	nodes := map[xy]*node{}
	xlen := len(fields)
	ylen := len(fields[0])
	endxy := xy{xlen*5 - 1, ylen*5 - 1}
	multi := 5

	// got nasty for p2
	for i := 0; i < multi; i++ {
		for j := 0; j < multi; j++ {
			for x, line := range fields {
				x = (xlen * i) + x
				for y, riskLevel := range line {
					y = (ylen * j) + y
					riskLevelInt, _ := strconv.Atoi(string(riskLevel))

					if i > 0 || j > 0 {
						riskLevelInt = (riskLevelInt + j + i) % 9

						if riskLevelInt == 0 {
							riskLevelInt = 9
						}
					}

					nodes[xy{x, y}] = &node{x, y, riskLevelInt, false, (endxy.x - x) + (endxy.y - y), 0, []*node{}}

					if i == multi-1 && j == multi-1 {
						if x == len(fields)*multi-1 && y == len(line)*multi-1 {
							nodes[xy{x, y}].terminus = true
						}
					}
				}
			}
		}
	}

	return nodes
}

func assignNeighbors(nodes map[xy]*node) {
	for k, node := range nodes {
		neighborPos := []xy{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

		for _, np := range neighborPos {
			nxy := xy{k.x + np.x, k.y + np.y}

			if neighborNode, ok := nodes[nxy]; ok {
				node.neighbors = append(node.neighbors, neighborNode)
			}
		}
	}
}

func findPath(nodes map[xy]*node) []*node {
	// A*-ish using pqueue/heap, and riskLevel+terminusDistance heuristic
	openSet := queue{}
	heap.Init(&openSet)
	heap.Push(&openSet, nodes[xy{0, 0}])
	gScore := map[xy]int{}
	cameFrom := map[xy]*node{}

	for openSet.Len() > 0 {
		curr := heap.Pop(&openSet).(*node)
		if curr.terminus {
			return build(curr, cameFrom)
		}

		for _, neighbor := range curr.neighbors {
			nxy := xy{neighbor.x, neighbor.y}
			currxy := xy{curr.x, curr.y}
			fScore := gScore[currxy] + neighbor.riskLevel

			if gScore[nxy] == 0 || fScore < gScore[nxy] {
				cameFrom[nxy] = curr
				gScore[nxy] = fScore
				nodes[nxy].score = fScore + nodes[nxy].terminusDistance
				heap.Push(&openSet, nodes[nxy])
			}
		}
	}

	return nil
}

func build(terminus *node, cameFrom map[xy]*node) []*node {
	history := []*node{}
	next := terminus
	for next.x >= 0 || next.y >= 0 {
		if next.x == 0 && next.y == 0 {
			break
		}
		history = append(history, next)
		next = cameFrom[xy{next.x, next.y}]
	}

	return history
}
