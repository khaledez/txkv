package kv

type MemoryKvStore struct {
	store             map[string]string
	transactionsStack Stack
}

func NewMemoryStore() *MemoryKvStore {
	return &MemoryKvStore{
		store:             make(map[string]string),
		transactionsStack: Stack{},
	}
}

func (s *MemoryKvStore) Get(key string) (string, bool) {
	if val, ok := s.store[key]; ok {
		return val, true
	}
	return "", false
}

func (s *MemoryKvStore) Set(key, value string) {
	s.store[key] = value
}

func (s *MemoryKvStore) Delete(key string) bool {
	if _, ok := s.store[key]; ok {
		delete(s.store, key)
		return true
	}

	return false
}

func (s *MemoryKvStore) Count(value string) int {
	return len(s.store)
}

func (s *MemoryKvStore) BeginTransaction() error {
	return nil
}

func (s *MemoryKvStore) CommitTransaction() error {
	return nil
}

func (s *MemoryKvStore) RollbackTransaction() error {
	return nil
}
