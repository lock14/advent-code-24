package deque

import (
	"fmt"
	"iter"
	"strings"
)

const (
	// DefaultCapacity is the capacity assigned if no other is provided.
	DefaultCapacity = 1
	// if an arraydeque's capacity is under this amount its capacity
	// will double when it needs to be resized.
	doublingThreshold = 512
	// if an arraydeque's capacity is under this amount its capacity
	// will increase by 50% when it needs to be resized.
	fiftyPercentThreshold = 2048
)

// Deque represents a deque of elements of type T backed by an array.
// The zero value for Deque is an empty deque ready to use.
type Deque[T any] struct {
	slice []T
	front int
	back  int
	size  int
}

// Config holds the values for configuring a Deque.
type Config struct {
	Capacity int
}

// Option configures a Deque config
type Option func(*Config)

// New creates an empty Deque whose initial size is 0.
func New[T any](opts ...Option) *Deque[T] {
	config := defaultConfig()
	for _, option := range opts {
		option(config)
	}
	return &Deque[T]{
		slice: make([]T, config.Capacity),
	}
}

// Add is an alias
func (d *Deque[T]) Add(t T) {
	d.AddBack(t)
}

// Remove is an alias
func (d *Deque[T]) Remove() T {
	return d.RemoveFront()
}

// Push is an alias
func (d *Deque[T]) Push(t T) {
	d.AddFront(t)
}

// Pop is an alias
func (d *Deque[T]) Pop() T {
	return d.RemoveFront()
}

func (d *Deque[T]) AddFront(t T) {
	if d.size == len(d.slice) {
		d.resize()
	}
	d.front--
	if d.front == -1 {
		d.front = len(d.slice) - 1
	}
	d.slice[d.front] = t
	d.size++
}

func (d *Deque[T]) RemoveFront() T {
	if d.Empty() {
		panic("cannot remove from an empty Deque")
	}
	var zero T
	t := d.slice[d.front]
	d.slice[d.front] = zero
	d.front++
	if d.front == len(d.slice) {
		d.front = 0
	}
	d.size--
	return t
}

func (d *Deque[T]) AddBack(t T) {
	if d.size == len(d.slice) {
		d.resize()
	}
	d.slice[d.back] = t
	d.back++
	if d.back == len(d.slice) {
		d.back = 0
	}
	d.size++
}

func (d *Deque[T]) RemoveBack() T {
	if d.Empty() {
		panic("cannot remove from an empty Deque")
	}
	var zero T
	d.back--
	if d.back == -1 {
		d.back = len(d.slice) - 1
	}
	t := d.slice[d.back]
	d.slice[d.back] = zero
	d.size--
	return t
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) Clear() {
	d.slice = nil
	d.front = 0
	d.back = 0
	d.size = 0
}

func (d *Deque[T]) String() string {
	str := make([]string, 0, d.Size())
	for t := range d.All() {
		str = append(str, fmt.Sprintf("%+v", t))
	}
	return "[" + strings.Join(str, ", ") + "]"
}

func (d *Deque[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for i := d.front; i < len(d.slice); i++ {
			if count == d.size || !yield(d.slice[i]) {
				return
			}
			count++
		}
		for i := 0; i < d.front; i++ {
			if count == d.size || !yield(d.slice[i]) {
				return
			}
			count++
		}
	}
}

func (d *Deque[T]) resize() {
	var newCap int
	if d.slice == nil {
		newCap = DefaultCapacity
	} else { // double
		newCap = len(d.slice)
		newCap += len(d.slice)
	}
	s := make([]T, newCap)
	m := copy(s, d.slice[d.front:])
	n := copy(s[m:], d.slice[0:d.front])
	if m+n != d.size {
		panic("resize algorithm incorrect")
	}
	d.slice = s
	d.front = 0
	d.back = d.size
}

func defaultConfig() *Config {
	return &Config{
		Capacity: DefaultCapacity,
	}
}