package application

import (
	"fmt"
	"sync"
)

var ctx Context[any]

type Context[T any] struct {
	ctx sync.Map
}

func init() {
	ctx = Context[any]{
		ctx: sync.Map{},
	}
}

func SetBean[T any](name string, bean T) {
	ctx.ctx.Store(name, bean)
}

func GetBean[T any](name string) (T, error) {
	value, ok := ctx.ctx.Load(name)
	if !ok {
		return value.(T), fmt.Errorf("bean %s not found", name)
	}
	return value.(T), nil
}
