package kv

type KeyValueStore interface {
	Get(key string) (string, bool)
	Set(key, value string)
	Delete(key string) bool
	Count(value string) int
	BeginTransaction() error
	CommitTransaction() error
	RollbackTransaction() error
}
