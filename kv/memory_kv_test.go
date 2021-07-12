package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_commit_transaction(t *testing.T) {
	kv := NewMemoryStore()

	kv.BeginTransaction()
	kv.Set("foo", "456")
	err := kv.CommitTransaction()
	assert.Nil(t, err)

	err = kv.RollbackTransaction()
	if assert.Error(t, err) {
		assert.Equal(t, "no transaction", err.Error())
	}

	val, ok := kv.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "456", val)
}

func Test_nested_transactions(t *testing.T) {
	// Given ..
	kv := NewMemoryStore()
	kv.Set("foo", "123")
	kv.Set("bar", "abc")
	kv.BeginTransaction()
	kv.Set("foo", "456")
	kv.BeginTransaction()
	kv.Set("foo", "789")

	val, ok := kv.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "789", val)

	kv.RollbackTransaction()

	val, ok = kv.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "456", val)

	kv.RollbackTransaction()

	val, ok = kv.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "123", val)
}
