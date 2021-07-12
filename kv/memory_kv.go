package kv

import "fmt"

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
	kv := s
	if currentTx, ok := s.transactionsStack.Peek(); ok {
		kv = currentTx.(*MemoryKvStore)
	}

	if val, ok := kv.store[key]; ok {
		return val, true
	}
	return "", false
}

func (s *MemoryKvStore) Set(key, value string) {
	kv := s
	if currentTx, ok := s.transactionsStack.Peek(); ok {
		kv = currentTx.(*MemoryKvStore)
	}

	kv.store[key] = value
}

func (s *MemoryKvStore) Delete(key string) bool {
	kv := s
	if currentTx, ok := s.transactionsStack.Peek(); ok {
		kv = currentTx.(*MemoryKvStore)
	}

	if _, ok := kv.store[key]; ok {
		delete(kv.store, key)
		return true
	}

	return false
}

func (s *MemoryKvStore) Count(value string) int {

	kv := s
	if currentTx, ok := s.transactionsStack.Peek(); ok {
		kv = currentTx.(*MemoryKvStore)
	}
	result := 0

	for _, v := range kv.store {
		if v == value {
			result++
		}
	}

	return result
}

func (s *MemoryKvStore) BeginTransaction() error {
	s.transactionsStack.Push(NewMemoryStore())
	return nil
}

func (s *MemoryKvStore) CommitTransaction() error {
	if txData, ok := s.transactionsStack.Pop(); ok {
		for k, v := range txData.(*MemoryKvStore).store {
			s.store[k] = v
		}
		return nil
	}
	return fmt.Errorf("no transaction")
}

func (s *MemoryKvStore) RollbackTransaction() error {
	if _, ok := s.transactionsStack.Pop(); ok {
		return nil
	}
	return fmt.Errorf("no transaction")
}
