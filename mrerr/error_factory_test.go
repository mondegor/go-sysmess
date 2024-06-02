package mrerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppErrorFactory_init(t *testing.T) {
	t.Parallel()

	calledGenerateID := false
	calledCaller := false

	factory := AppErrorFactory{
		generateID: func() string {
			calledGenerateID = true
			return ""
		},
		caller: func() StackTracer {
			calledCaller = true
			return nil
		},
	}
	item := AppError{}

	factory.init(&item)

	assert.True(t, calledGenerateID)
	assert.True(t, calledCaller)
}

func TestAppErrorFactory_initWithoutFunctions(t *testing.T) {
	t.Parallel()

	factory := AppErrorFactory{}
	item := AppError{}

	assert.NotPanics(t, func() { factory.init(&item) })
}
