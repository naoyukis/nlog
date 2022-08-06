package nlog

import (
	"io"
	"sync"
)

type Logger struct {
	mu        sync.Mutex
	prefix    string
	flag      int
	out       io.Writer
	buf       []byte
	isDiscard int32
}
