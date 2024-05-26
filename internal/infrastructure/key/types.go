package key

import (
	"fmt"
	"time"
)

type (
	Key struct {
		Tag   string
		Value string
	}

	KeyValue struct {
		Key   Key
		Value string
		Exp   time.Duration
	}

	ValueEntity interface {
		GetKV() KeyValue
		Scan(key Key, value string) error
	}
)

func (k *Key) Build() string {
	return fmt.Sprintf("%s-%s", k.Tag, k.Value)
}
