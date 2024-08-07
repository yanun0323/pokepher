package utils

import (
	"sync/atomic"
)

type Value[T any] struct {
	defaultValue T
	value        atomic.Value
}

func NewValue[T any](defaultValue T) *Value[T] {
	return &Value[T]{
		defaultValue: defaultValue,
		value:        atomic.Value{},
	}
}

func (v *Value[T]) Load() T {
	if v.value.Load() == nil {
		return v.defaultValue
	}

	return v.value.Load().(T)
}

func (v *Value[T]) Store(value T) {
	v.value.Store(value)
}
