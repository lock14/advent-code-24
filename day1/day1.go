package day1

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	list1 := &Int64Heap{}
	list2 := &Int64Heap{}
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		a, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		list1.Push(a)
		list2.Push(b)
	}
	heap.Init(list1)
	heap.Init(list2)
	distance := int64(0)
	for list1.Len() > 0 {
		a := heap.Pop(list1).(int64)
		b := heap.Pop(list2).(int64)
		distance += diff(a, b)
	}
	return distance
}

func Part2(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	list1 := make([]int64, 0)
	list2Counts := make(map[int64]int64)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		a, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, a)
		if _, ok := list2Counts[b]; !ok {
			list2Counts[b] = 0
		}
		list2Counts[b] += 1
	}
	similarity := int64(0)
	for _, a := range list1 {
		similarity += a * list2Counts[a]
	}
	return similarity
}

func diff(a int64, b int64) int64 {
	if a < b {
		return b - a
	}
	return a - b
}

// Int64Heap is a min-heap of int64.
type Int64Heap []int64

func (h *Int64Heap) Len() int           { return len(*h) }
func (h *Int64Heap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *Int64Heap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *Int64Heap) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *Int64Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
