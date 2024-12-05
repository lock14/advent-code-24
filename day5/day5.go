package day5

import (
	. "advent/util"
	"advent/util/deque"
	"advent/util/set"
	"strings"
)

func Part1(filename string) int64 {
	numbers := make(map[int64]*set.Set[int64])
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		parts := strings.Split(scanner.Text(), "|")
		a := Must(ParseInt64(parts[0]))
		b := Must(ParseInt64(parts[1]))
		if numbers[a] == nil {
			numbers[a] = set.New[int64]()
		}
		numbers[a].Add(b)
	}
	result := int64(0)
	for scanner.Scan() {
		list := Must(MapSliceErr(strings.Split(scanner.Text(), ","), ParseInt64))
		for _, n := range list {
			if numbers[n] == nil {
				numbers[n] = set.New[int64]()
			}
		}
		if topSort(numbers, list) {
			result += list[len(list)/2]
		}
	}
	return result
}

func Part2(filename string) int64 {
	numbers := make(map[int64]*set.Set[int64])
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		parts := strings.Split(scanner.Text(), "|")
		a := Must(ParseInt64(parts[0]))
		b := Must(ParseInt64(parts[1]))
		if numbers[a] == nil {
			numbers[a] = set.New[int64]()
		}
		numbers[a].Add(b)
	}
	result := int64(0)
	for scanner.Scan() {
		list := Must(MapSliceErr(strings.Split(scanner.Text(), ","), ParseInt64))
		for _, n := range list {
			if numbers[n] == nil {
				numbers[n] = set.New[int64]()
			}
		}
		if !topSort(numbers, list) {
			result += list[len(list)/2]
		}
	}
	return result
}

func topSort(numbers map[int64]*set.Set[int64], list []int64) bool {
	visiting := set.New[int64]()
	visited := set.New[int64]()
	stack := deque.New[int64]()
	filteredVertices := set.New[int64]()
	filteredVertices.Add(list...)
	filter := func(n int64) bool { return filteredVertices.Contains(n) }
	for _, n := range list {
		dfs(n, numbers, filter, visiting, visited, stack)
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

func dfs(v int64, graph map[int64]*set.Set[int64], test func(int64) bool, visiting, visited *set.Set[int64], stack *deque.Deque[int64]) bool {
	if visited.Contains(v) {
		return false
	}
	if visiting.Contains(v) {
		// contains a cycle
		return true
	}
	visiting.Add(v)
	for n := range graph[v].All() {
		if test(n) && dfs(n, graph, test, visiting, visited, stack) {
			return true
		}
	}
	visiting.Remove(v)
	visited.Add(v)
	stack.Push(v)
	return false
}
