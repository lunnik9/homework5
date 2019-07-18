package main

type Strings []string

func (s *Strings) Len() int {
	return len(*s)
}

func (s *Strings) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Strings) Less(i, j int) bool {
	minL := 0
	if len((*s)[i]) <= len((*s)[j]) {
		minL = len((*s)[i])
	} else {
		minL = len((*s)[j])
	}
	for ii := 0; ii < minL; ii++ {
		if (*s)[i][ii] == (*s)[j][ii] {
			continue
		} else if (*s)[i][ii] < (*s)[j][ii] {
			return true
		}
	}
	if minL == len((*s)[i]) {
		return true
	}
	return false
}

func (s *Strings) BubbleSort() {
	for i := 0; i < s.Len(); i++ {
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(i, j) {
				s.Swap(i, j)
			}
		}
	}
}
