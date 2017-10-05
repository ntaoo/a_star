package a_star

type Node struct {
	h              int // heuristic distance
	f              int
	g              int
	inherentCost   int
	parent         *Node
	connectedNodes []*Node
	isInOpenSet    bool
	isInClosedSet  bool
}

func NewNode(inherentCost int) *Node {
	return &Node{inherentCost: inherentCost}
}

const InfiniteDistance int = 4294967295

type Graph struct {
	allNodes []*Node
}

type GraphTrait interface {
	GetAllNodes() []*Node
	GetDistance(a *Node, b *Node) int
	GetHeuristicDistance(a *Node, b *Node) int
	GetNeighboursOf(node *Node) []*Node
}

func AStar(graph GraphTrait, start *Node, goal *Node) []*Node {
	allNodes := graph.GetAllNodes()
	for _, node := range allNodes {
		node.isInClosedSet = false
		node.isInOpenSet = false
		node.parent = nil
	}

	open := []*Node{}
	var lastClosed *Node

	open = append(open, start)
	start.isInOpenSet = true
	start.f = -1
	start.g = -1

	for len(open) != 0 {
		currentNode := findLowestCostNode(open)

		if currentNode == goal {
			return buildSolvedPath(goal)
		}

		currentNode.isInOpenSet = false
		currentNode.isInClosedSet = true
		lastClosed = currentNode
		open = deleteNode(currentNode, open)

		neighbors := graph.GetNeighboursOf(currentNode)
		for _, candidate := range neighbors {
			distance := graph.GetDistance(currentNode, candidate)
			if distance != InfiniteDistance || candidate == goal {
				if candidate.isInClosedSet {
					continue
				}

				if !candidate.isInOpenSet {
					candidate.parent = lastClosed

					candidate.g = currentNode.g + distance
					h := graph.GetHeuristicDistance(candidate, goal)
					candidate.f = candidate.g + h

					candidate.isInOpenSet = true
					open = append(open, candidate)
				}
			}
		}
	}
	// No path found.
	return []*Node{}
}
func buildSolvedPath(goal *Node) []*Node {
	path := make([]*Node, 0)
	path = append(path, goal)
	currentNode := path[0]
	for currentNode.parent != nil {
		currentNode = currentNode.parent
		path = append([]*Node{currentNode}, path...)
	}
	return path
}

// Find a lowest cost node.
// The nodes len MUST be grater than equals to 1.
func findLowestCostNode(nodes []*Node) *Node {
	if len(nodes) == 1 {
		return nodes[0]
	}

	retNode := nodes[0]
	for _, node := range nodes[1:] {
		if retNode.f >= node.f {
			retNode = node
		}
	}

	return retNode
}

func deleteNode(node *Node, nodes []*Node) []*Node {
	for i, n := range nodes {
		if n == node {
			nodes = append(nodes[:i], nodes[i+1:]...)
			break
		}
	}
	return nodes
}
