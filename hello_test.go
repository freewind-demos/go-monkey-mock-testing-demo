package main

import (
	"testing"
	"github.com/bouk/monkey"
	"time"
	"github.com/magiconair/properties/assert"
	"reflect"
)

func TestMockFunction(t *testing.T) {
	patches := monkey.Patch(Hello, func() string {
		return "Hello from test"
	})
	defer patches.Unpatch()
	assert.Equal(t, Hello(), "Hello from test")
}

func TestMockMethod(t *testing.T) {
	clock := &Clock{
		name: "testing-clock",
	}
	patches := monkey.PatchInstanceMethod(reflect.TypeOf(clock), "GetNow", func(_ *Clock) time.Time {
		return time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC)
	})
	defer patches.Unpatch()
	assert.Equal(t, clock.GetNow().Day(), 11)
}
