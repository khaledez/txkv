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

type ActionID uint8

const (
	SetOp ActionID = iota
	DeleteOp
)

type WriteLog struct {
	Op    ActionID
	Key   string
	Value string
}

type StringSet map[string]struct{}

func (s *StringSet) Append(key string) {
	(*s)[key] = struct{}{}
}

func (s StringSet) Contains(key string) bool {
	_, ok := s[key]
	return ok
}
