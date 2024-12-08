package day7

import (
	. "advent/util"
	"advent/util/deque"
	"strings"
)

var ops = map[uint8]func(int64, int64) int64{
	'+': func(a int64, b int64) int64 {
		return a + b
	},
	'*': func(a int64, b int64) int64 {
		return a * b
	},
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
		if equationHolds(deq, target) {
			sum += target
		}
		deq.Clear()
	}
	return sum
}

func Part2(filename string) int64 {
	return 0
}

func equationHolds(deq *deque.Deque[int64], target int64) bool {
	if deq.Size() == 1 && deq.PeekFront() == target {
		return true
	} else if deq.Size() > 1 {
		a := deq.RemoveFront()
		b := deq.RemoveFront()
		for _, op := range ops {
			deq.AddFront(op(a, b))
			if equationHolds(deq, target) {
				return true
			}
			deq.RemoveFront()
		}
		deq.AddFront(b)
		deq.AddFront(a)
	}
	return false
}
