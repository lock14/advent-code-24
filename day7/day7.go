package day7

import (
	. "advent/util"
	"advent/util/deque"
	"math"
	"strings"
)

type Op func(int64, int64) int64

func add(a int64, b int64) int64 {
	return a + b
}

func multiply(a int64, b int64) int64 {
	return a * b
}

func concat(a int64, b int64) int64 {
	numDigits := math.Floor(math.Log10(float64(b))) + 1.0
	r := (a * int64(math.Pow(10, numDigits))) + b
	return r
}

func Part1(filename string) int64 {
	var sum int64
	deq := deque.New[int64]()
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		target := Must(ParseInt64(parts[0]))
		deq.AddBack(Must(MapSliceErr(strings.Fields(parts[1]), ParseInt64))...)
		if equationHolds(deq, target, []Op{add, multiply}) {
			sum += target
		}
		deq.Clear()
	}
	return sum
}

func Part2(filename string) int64 {
	var sum int64
	deq := deque.New[int64]()
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		target := Must(ParseInt64(parts[0]))
		deq.AddBack(Must(MapSliceErr(strings.Fields(parts[1]), ParseInt64))...)
		if equationHolds(deq, target, []Op{add, multiply, concat}) {
			sum += target
		}
		deq.Clear()
	}
	return sum
}

func equationHolds(stack *deque.Deque[int64], target int64, ops []Op) bool {
	if stack.Size() == 1 && stack.Peek() == target {
		return true
	} else if stack.Size() > 1 {
		a := stack.Pop()
		b := stack.Pop()
		for _, op := range ops {
			stack.Push(op(a, b))
			if equationHolds(stack, target, ops) {
				return true
			}
			stack.Pop()
		}
		stack.Push(b)
		stack.Push(a)
	}
	return false
}
