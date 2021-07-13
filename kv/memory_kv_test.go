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

func Test_reading_from_parent_transaction_if_not_exist(t *testing.T) {
	kv := NewMemoryStore()
	kv.Set("foo", "123")
	kv.BeginTransaction()
	kv.Set("bar", "456")

	// 1. read from parent
	val, ok := kv.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "123", val)

	// 2. read after rollback
	val, ok = kv.Get("bar")
	assert.True(t, ok)
	assert.Equal(t, "456", val)
	kv.RollbackTransaction()
	_, ok = kv.Get("bar")
	assert.False(t, ok)
}

func Test_count_should_not_include_parent_value_if_key_is_modified_in_current_tx(t *testing.T) {
	kv := NewMemoryStore()

	assert.Zero(t, kv.Count("123"))

	kv.Set("foo", "123")
	kv.Set("bar", "123")
	assert.Equal(t, 2, kv.Count("123"))

	kv.BeginTransaction()
	kv.Set("foo", "456")
	kv.Set("number", "123")

	assert.Equal(t, 2, kv.Count("123"))
}
