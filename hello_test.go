package main

import (
	"testing"
	"github.com/bouk/monkey"
	"time"
	"github.com/magiconair/properties/assert"
)

func TestDayOfMonth(t *testing.T) {
	patches := monkey.Patch(GetNow, func() time.Time {
		return time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC)
	})
	defer patches.Unpatch()
	assert.Equal(t, DayOfMonth(), 11)
}
