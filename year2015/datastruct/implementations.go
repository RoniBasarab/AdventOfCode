package datastruct

type Stack []interface{}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return 0
	}

	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *Stack) Peek() interface{} {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
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

type Set map[interface{}]interface{}

func (s Set) AddAndReturnCurrentSantaPosition(item Coordinates2D, uniqueHouses *int) Coordinates2D {
	if !(s.IsInSet(item)) {
		s[item] = Coordinates2D{X: item.X, Y: item.Y}
		*uniqueHouses++
		return item
	}
	return item
}

func (s Set) Add(item interface{}) {
	s[item] = struct{}{}
}

func (s Set) Remove(item interface{}) {
	delete(s, item)
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) IsInSet(item interface{}) bool {
	_, ok := s[item]
	return ok

}

func (s Set) SymmetricDifference(other *Set) Set {
	symmetricDifferenceSet := Set{}

	for _, setItem := range s {
		symmetricDifferenceSet.Add(setItem)
	}

	for _, setItem := range *other {
		symmetricDifferenceSet.Add(setItem)
	}

	return symmetricDifferenceSet
}
