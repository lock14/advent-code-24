package graph

import (
	"advent/util/set"
	"iter"
)

type Graph[V comparable] struct {
	directed bool
	g        map[V]*set.Set[V]
}

type Opt[V comparable] func(graph *Graph[V])

func Directed[V comparable]() Opt[V] {
	return func(g *Graph[V]) {
		g.directed = true
	}
}

func New[V comparable](opts ...Opt[V]) *Graph[V] {
	g := &Graph[V]{
		g:        make(map[V]*set.Set[V]),
		directed: false,
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func (g *Graph[V]) Directed() bool {
	return g.directed
}

func (g *Graph[V]) ContainsVertex(vertex V) bool {
	_, ok := g.g[vertex]
	return ok
}

func (g *Graph[V]) AddVertex(vertices ...V) {
	for _, vertex := range vertices {
		if !g.ContainsVertex(vertex) {
			g.g[vertex] = set.New[V]()
		}
	}
}

func (g *Graph[V]) RemoveVertex(vertex V) {
	delete(g.g, vertex)
}

func (g *Graph[V]) ContainsEdge(u, v V) bool {
	return g.ContainsVertex(u) && g.ContainsVertex(v) && g.g[u].Contains(v)
}

func (g *Graph[V]) AddEdge(vertices ...V) {
	for i := 1; i < len(vertices); i += 2 {
		g.AddVertex(vertices[i-1])
		g.AddVertex(vertices[i])
		g.g[vertices[i-1]].Add(vertices[i])
		if !g.directed {
			g.g[vertices[i]].Add(vertices[i-1])
		}
	}
}

func (g *Graph[V]) RemoveEdge(u, v V) {
	if g.ContainsEdge(u, v) {
		g.g[u].Remove(v)
		if !g.directed {
			g.g[v].Remove(u)
		}
	}
}

func (g *Graph[V]) Vertices() iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range g.g {
			if !yield(v) {
				return
			}
		}
	}
}

func (g *Graph[V]) Neighbors(v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		if g.ContainsVertex(v) {
			for neighbor := range g.g[v].All() {
				if !yield(neighbor) {
					return
				}
			}
		}
	}
}

func (g *Graph[V]) Edges() iter.Seq2[V, V] {
	shouldYield := shouldYieldFunc(g)
	return func(yield func(V, V) bool) {
		for u := range g.Vertices() {
			for v := range g.Neighbors(u) {
				if shouldYield(u, v) && !yield(u, v) {
					return
				}
			}
		}
	}
}

func shouldYieldFunc[V comparable](g *Graph[V]) func(V, V) bool {
	if g.directed {
		return func(v V, v2 V) bool {
			return true
		}
	} else {
		cache := make(map[V]map[V]struct{})
		return func(u V, v V) bool {
			if _, ok := cache[u]; !ok {
				cache[u] = make(map[V]struct{})
			}
			if _, ok := cache[u][v]; ok {
				return false
			}
			if _, ok := cache[v][u]; ok {
				return false
			}
			cache[u][v] = struct{}{}
			cache[v][u] = struct{}{}
			return true
		}
	}
}
