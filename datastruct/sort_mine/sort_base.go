package sort_mine

type SortStrategy interface {
	Sort(arr []int)
}

type Sorter struct {
	strategy SortStrategy
	arr      []int
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Execute() {
	s.strategy.Sort(s.arr)
}
