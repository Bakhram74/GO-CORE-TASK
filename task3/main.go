package main

// tsk3
type Node struct {
	Key   string
	Value int
	Next  *Node
}
type HashMap struct {
	buckets []*Node
	size    int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func hashFunction(key string, size int) uint {
	return uint(size % len(key))
}

func (hm *HashMap) Add(key string, value int) {
	node := &Node{Key: key, Value: value}
	index := hashFunction(key, hm.size)

	if hm.buckets[index] == nil {
		hm.buckets[index] = node
	} else {
		current := hm.buckets[index]
		for current != nil {

			if current.Key == key {
				current.Value = value
				return
			}
			if current.Next == nil {
				break
			}
			current = current.Next
		}
		current.Next = node
	}
}

func (hm *HashMap) Get(key string) (int, bool) {
	index := hashFunction(key, hm.size)
	current := hm.buckets[index]

	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	return 0, false
}

func (hm *HashMap) Copy() map[string]int {
	newMap := map[string]int{}

	for i := range hm.buckets {
		current := hm.buckets[i]
		for current != nil {
			newMap[current.Key] = current.Value
			current = current.Next
		}
	}
	return newMap
}

func (hm *HashMap) Exists(key string) bool {
	index := hashFunction(key, hm.size)
	current := hm.buckets[index]

	for current != nil {
		if current.Key == key {
			return true
		}
		current = current.Next
	}
	return false
}

func (hm *HashMap) Remove(key string) {
	index := hashFunction(key, hm.size)
	current := hm.buckets[index]
	var prev *Node

	for current != nil {

		if current.Key == key {

			if prev == nil {
				hm.buckets[index] = current.Next
			} else {

				prev.Next = current.Next
			}

			return
		}
		prev = current
		current = current.Next
	}
}
