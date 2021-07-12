package kv

// Stack is used to manage transactions in the KV store by storing a subset of store per transaction
// this implemention is intentional not thread-safe.
type Stack []KeyValueStore

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(entry KeyValueStore) {
	*s = append(*s, entry)
}

func (s *Stack) Pop() (KeyValueStore, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return element, true
}

func (s *Stack) Peek() (KeyValueStore, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	return (*s)[len(*s)-1], true
}
