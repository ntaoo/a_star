package a_star

import (
	"testing"
)

// Sample Graph for the test.
type SampleGraph struct {
	Graph
	start *Node
	a     *Node
	b     *Node
	c     *Node
	d     *Node
	goal  *Node
}

func NewSampleGraph() *SampleGraph {
	start := NewNode(1)
	a := NewNode(2)
	b := NewNode(1)
	c := NewNode(1)
	d := NewNode(2)
	goal := NewNode(3)

	start.connectedNodes = []*Node{a, b}
	a.connectedNodes = []*Node{start, b, c, d}
	b.connectedNodes = []*Node{start, a, c, d}
	c.connectedNodes = []*Node{a, b, d, goal}
	d.connectedNodes = []*Node{a, b, c, goal}
	goal.connectedNodes = []*Node{c, d}
	return &SampleGraph{start: start, a: a, b: b, c: c, d: d, goal: goal}
}

func (graph *SampleGraph) GetAllNodes() []*Node {
	return []*Node{graph.start, graph.a, graph.b, graph.c, graph.d, graph.goal}
}

func (graph *SampleGraph) GetDistance(a *Node, b *Node) int {
	return b.inherentCost
}

func (graph *SampleGraph) GetHeuristicDistance(node *Node, goalNode *Node) int {
	if node == graph.start {
		return 3
	}
	if node == graph.a || node == graph.b {
		return 2
	}
	if node == graph.c || node == graph.d {
		return 1
	}
	return 0
}

func (graph *SampleGraph) GetNeighboursOf(node *Node) []*Node {
	return node.connectedNodes
}

func TestAStarFindsFastestPath(t *testing.T) {
	graph := NewSampleGraph()

	actualPath := AStar(graph, graph.start, graph.goal)
	actualLen := len(actualPath)
	expectedLen := 4

	if actualLen != expectedLen {
		t.Errorf("expected solved path len is %v, actual: %v", expectedLen, actualLen)
	}
	if actualPath[0] != graph.start || actualPath[1] != graph.b || actualPath[2] != graph.c || actualPath[3] != graph.goal {
		t.Errorf("unexpected solved path.")
	}
}

func TestAStarDoesNotFindWithImpassibleNodes(t *testing.T) {
	graph := NewSampleGraph()

	graph.c.inherentCost = InfiniteDistance
	graph.d.inherentCost = InfiniteDistance

	actualPath := AStar(graph, graph.start, graph.goal)
	actualLen := len(actualPath)
	expectedLen := 0

	if actualLen != expectedLen {
		t.Errorf("expected solved path len is %v (means 'no path found'), actual: %v", expectedLen, actualLen)
	}
}
