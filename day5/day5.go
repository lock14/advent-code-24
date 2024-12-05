package day5

import (
	. "advent/util"
	"advent/util/deque"
	"advent/util/graph"
	"advent/util/set"
	"strings"
)

func Part1(filename string) int64 {
	result := int64(0)
	ScanAndProcess(filename, func(n int64) {
		result += n
	}, func(int64) {})
	return result
}

func Part2(filename string) int64 {
	result := int64(0)
	ScanAndProcess(filename, func(int64) {}, func(n int64) {
		result += n
	})
	return result
}

func ScanAndProcess(filename string, correctOrder func(int64), incorrectOrder func(int64)) {
	numbers := graph.New(graph.Directed[int64]())
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		parts := strings.Split(scanner.Text(), "|")
		a := Must(ParseInt64(parts[0]))
		b := Must(ParseInt64(parts[1]))
		numbers.AddEdge(a, b)
	}
	for scanner.Scan() {
		list := Must(MapSliceErr(strings.Split(scanner.Text(), ","), ParseInt64))
		if topSort(numbers, list) {
			correctOrder(list[len(list)/2])
		} else {
			incorrectOrder(list[len(list)/2])
		}
	}
}

func topSort(graph *graph.Graph[int64], list []int64) bool {
	visiting := set.New[int64]()
	visited := set.New[int64]()
	stack := deque.New[int64]()
	listVertices := set.New[int64]()
	listVertices.Add(list...)
	shouldRecurse := func(n int64) bool { return listVertices.Contains(n) }
	for _, n := range list {
		dfs(n, graph, shouldRecurse, visiting, visited, stack)
	}
	i := 0
	correctOrder := true
	for n := range stack.All() {
		if list[i] != n {
			list[i] = n
			correctOrder = false
		}
		i++
	}
	return correctOrder
}

func dfs(v int64, graph *graph.Graph[int64], shouldRecurse func(int64) bool, visiting, visited *set.Set[int64], stack *deque.Deque[int64]) bool {
	if visited.Contains(v) {
		// already processed v
		return false
	}
	if visiting.Contains(v) {
		// graph contains a cycle
		return true
	}
	visiting.Add(v)
	for n := range graph.Neighbors(v) {
		if shouldRecurse(n) && dfs(n, graph, shouldRecurse, visiting, visited, stack) {
			// pass cycle detection up
			return true
		}
	}
	visiting.Remove(v)
	visited.Add(v)
	stack.Push(v)
	return false
}
