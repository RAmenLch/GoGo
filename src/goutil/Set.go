package goutil

type Set struct {
	items map[int]interface{}
}

func NewSet() *Set {
	set := &Set{items: make(map[int]interface{})}
	return set
}

type Item interface {
	Hash() int
}

func (s *Set) Add(item Item) {
	s.items[item.Hash()] = item
}

func (s *Set) Contains(item Item) bool {
	_, co := s.items[item.Hash()]
	return co
}

func (s *Set) Delete(items Item) {
	delete(s.items, items.Hash())
}

func (s *Set) Size() int {
	return len(s.items)
}
func (s *Set) Values() []interface{} {
	values := make([]interface{}, s.Size())
	count := 0
	for _, value := range s.items {
		values[count] = value
		count++
	}
	return values
}

func (s *Set) Hash() int {
	hash := 0
	for key, _ := range s.items {
		keyR := key % (1 << 8)
		hash = (hash << uint(5)) - hash + keyR
	}
	return hash
}
