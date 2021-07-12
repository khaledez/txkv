package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pushing_to_empty_stack_then_popping_keeps_it_empty(t *testing.T) {
	stack := Stack{}

	assert.True(t, stack.IsEmpty())

	stack.Push(&dummyKeyValueStore{"test"})
	assert.False(t, stack.IsEmpty())

	_, ok := stack.Pop()
	assert.True(t, ok)
	assert.True(t, stack.IsEmpty())
}

type dummyKeyValueStore struct {
	val string
}

func (s *dummyKeyValueStore) Get(key string) (string, bool) {
	if len(s.val) == 0 {
		return "", false
	}
	return s.val, true
}

func (s *dummyKeyValueStore) Set(key, value string) {
	s.val = value
}

func (s *dummyKeyValueStore) Delete(key string) bool {
	if len(s.val) == 0 {
		return false
	}
	s.val = ""

	return true
}

func (s *dummyKeyValueStore) Count(value string) int {
	return 1
}

func (s *dummyKeyValueStore) BeginTransaction() error {
	return nil
}

func (s *dummyKeyValueStore) CommitTransaction() error {
	return nil
}

func (s *dummyKeyValueStore) RollbackTransaction() error {
	return nil
}
