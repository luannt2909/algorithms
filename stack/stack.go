package stack

type Stack struct {
	Items []interface{}
}

func (s *Stack) Push(item interface{}) *Stack {
	s.Items = append(s.Items, item)
	return s
}

func (s *Stack) Len() int {
	return len(s.Items)
}

func (s *Stack) Empty() bool {
	return len(s.Items) <= 0
}

func (s *Stack) Pop() interface{} {
	if s.Len() <= 0 {
		return nil
	}
	item := s.Items[s.Len()-1]
	s.Items = s.Items[:s.Len()-1]
	return item
}
