package datastruct

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

type Coordinates2D struct {
	X int
	Y int
}

func (c *Coordinates2D) Plus(other *Coordinates2D) {
	c.X += (*other).X
	c.Y += (*other).Y
}

type Coordinates3D struct {
	X int
	Y int
	Z int
}

type Set struct {
	Items []interface{}
}

func (s *Set) AddAndReturnCurrentSantaPosition(item Coordinates2D, uniqueHouses *int) Coordinates2D {
	if !(s.IsInSet(item)) {
		s.Items = append(s.Items, item)
		*uniqueHouses++
		return item
	}
	return item
}

func (s *Set) Add(item interface{}) {
	if !(s.IsInSet(item)) {
		s.Items = append(s.Items, item)
	}
}

func (s *Set) Remove(item interface{}) {
	indexToRemove := -1

	for i, setItem := range s.Items {
		if setItem == item {
			indexToRemove = i
			break
		}
	}

	if indexToRemove != -1 {
		s.Items = append(s.Items[:indexToRemove], s.Items[indexToRemove+1:]...)
	}
}

func (s *Set) Size() int {
	return len(s.Items)
}

func (s *Set) IsInSet(item interface{}) bool {
	for _, setItem := range s.Items {
		if setItem == item {
			return true
		}
	}
	return false
}
