package kv

import "fmt"

type MemoryKvStore struct {
	store        map[string]string
	wal          []WriteLog
	affectedKeys StringSet
	parentTx     *MemoryKvStore
	currentTx    *MemoryKvStore
}

func NewMemoryStore() *MemoryKvStore {
	kv := &MemoryKvStore{
		store: make(map[string]string),
	}
	kv.currentTx = kv

	return kv
}

func newMemoryStore(parent *MemoryKvStore) *MemoryKvStore {
	return &MemoryKvStore{
		store:        make(map[string]string),
		wal:          make([]WriteLog, 0),
		affectedKeys: make(StringSet),
		parentTx:     parent,
	}
}

func (s *MemoryKvStore) Get(key string) (string, bool) {
	// 1. look for the key in the current transaction
	if val, ok := s.currentTx.store[key]; ok {
		return val, true
	}

	// 2. if not found check parent transactions
	if s.currentTx.parentTx != nil {
		cursor := s.currentTx.parentTx
		for cursor != nil {
			if val, ok := cursor.store[key]; ok {
				return val, true
			}
			cursor = cursor.parentTx
		}
	}

	// 3. not found any where? return false
	return "", false
}

func (s *MemoryKvStore) Count(value string) int {
	result := 0
	kv := s.currentTx
	// 1. count in the current transaction
	for _, val := range kv.store {
		if val == value {
			result++
		}
	}

	// 2. count in the parent transaction
	if kv.parentTx != nil {
		cursor := kv.parentTx
		for cursor != nil {
			for key, val := range cursor.store {
				if val == value && !kv.affectedKeys.Contains(key) {
					result++
				}
			}
			cursor = cursor.parentTx
		}
	}

	return result
}

func (s *MemoryKvStore) Set(key, value string) {
	kv := s.currentTx
	if kv.parentTx != nil {
		kv.wal = append(s.wal, WriteLog{Op: SetOp, Key: key, Value: value})
		kv.affectedKeys.Append(key)
	}

	kv.store[key] = value
}

func (s *MemoryKvStore) Delete(key string) bool {
	kv := s.currentTx
	if kv.parentTx != nil {
		kv.wal = append(s.wal, WriteLog{Op: DeleteOp, Key: key})
		kv.affectedKeys.Append(key)
	}

	if _, ok := kv.store[key]; ok {
		delete(kv.store, key)
		return true
	}

	return false
}

func (s *MemoryKvStore) BeginTransaction() error {
	newTx := newMemoryStore(s.currentTx)
	s.currentTx = newTx
	return nil
}

func (s *MemoryKvStore) CommitTransaction() error {
	kv := s.currentTx
	if kv.parentTx != nil {
		targetStore := kv.parentTx
		// if we are in a nested transaction, then we need to preserve WAL to be executed in the parent
		if targetStore.parentTx != nil {
			targetStore.wal = append(targetStore.wal, kv.wal...)
		}

		for _, action := range kv.wal {
			switch action.Op {
			case DeleteOp:
				delete(targetStore.store, action.Key)
			case SetOp:
				targetStore.store[action.Key] = action.Value
			}
		}

		s.currentTx = targetStore
		return nil
	}
	return fmt.Errorf("no transaction")
}

func (s *MemoryKvStore) RollbackTransaction() error {
	if s.currentTx.parentTx != nil {
		s.currentTx = s.currentTx.parentTx
		return nil
	}
	return fmt.Errorf("no transaction")
}
