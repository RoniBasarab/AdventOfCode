package datastruct

import "fmt"

type Stack struct {
	elements []rune
}

func (s *Stack) Pop() rune {
	if len(s.elements) == 0 {
		return 0
	}

	x := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return x
}

func (s *Stack) Peek() rune {
	if len(s.elements) == 0 {
		return 0
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Push(x rune) {
	s.elements = append(s.elements, x)
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

type FourWayNode struct {
	Up    *FourWayNode
	Down  *FourWayNode
	Left  *FourWayNode
	Right *FourWayNode
}

func (currentNode *FourWayNode) AppendAndReturnCurrentPosition(direction rune, uniqueHouses *int) *FourWayNode {
	newNode := &FourWayNode{Up: nil, Down: nil, Left: nil, Right: nil}

	switch direction {
	case '^':
		if currentNode.Up != nil {
			fmt.Println("I have someone up!, i go UP!")
			return currentNode.Up
		}
		fmt.Println("There's nobody up!, i create UP!")
		newNode.Down = currentNode
		currentNode.Up = newNode
		checkIfNeighboursExistAndConnect(&newNode)
		*uniqueHouses += 1
		return newNode

	case 'v':
		if currentNode.Down != nil {
			fmt.Println("I have someone down!, i go DOWN!")
			return currentNode.Down
		}
		fmt.Println("There's nobody down!, i create DOWN!")
		newNode.Up = currentNode
		currentNode.Down = newNode
		checkIfNeighboursExistAndConnect(&newNode)
		*uniqueHouses += 1
		return newNode

	case '>':
		if currentNode.Right != nil {
			fmt.Println("I have someone right!, i go Right!")
			return currentNode.Right
		}
		fmt.Println("There's nobody right!, i create Right!")
		newNode.Left = currentNode
		currentNode.Right = newNode
		checkIfNeighboursExistAndConnect(&newNode)
		*uniqueHouses += 1
		return newNode

	case '<':
		if currentNode.Left != nil {
			fmt.Println("I have someone left!, i go LEFT!")
			return currentNode.Left
		}
		fmt.Println("There's nobody left!, i create LEFT!")
		newNode.Right = currentNode
		currentNode.Left = newNode
		checkIfNeighboursExistAndConnect(&newNode)
		*uniqueHouses += 1
		return newNode

	default:
		return currentNode
	}
}

func checkIfNeighboursExistAndConnect(node **FourWayNode) {
	if (*node).Up != nil {
		(*node).Up.Down = *node
	}
	if (*node).Down != nil {
		(*node).Down.Up = *node
	}
	if (*node).Right != nil {
		(*node).Right.Left = *node
	}
	if (*node).Left != nil {
		(*node).Left.Right = *node
	}
}
