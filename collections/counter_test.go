package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_MostCommon(t *testing.T) {
	t.Run("NoValues", func(t *testing.T) {
		counter := NewCounter()
		assert.Empty(t, counter.MostCommon(5))
	})

	t.Run("NGreaterThanValues", func(t *testing.T) {
		counter := NewCounter()
		counter.AddItems("foo", "foo", "foo", "baz", "bar", "bar")
		assert.Equal(t, []CounterItem{
			{"foo", 3},
			{"bar", 2},
			{"baz", 1},
		}, counter.MostCommon(5))
	})

	t.Run("NLessThanValues", func(t *testing.T) {
		counter := NewCounter()
		counter.AddItems("foo", "foo", "foo", "baz", "bar", "bar")
		assert.Equal(t, []CounterItem{
			{"foo", 3},
		}, counter.MostCommon(1))
	})

	t.Run("NEqualToValues", func(t *testing.T) {
		counter := NewCounter()
		counter.AddItems("foo", "foo", "foo", "baz", "bar", "bar")
		assert.Equal(t, []CounterItem{
			{"foo", 3},
			{"bar", 2},
			{"baz", 1},
		}, counter.MostCommon(3))
	})

	t.Run("ValuesExistAfterGetting", func(t *testing.T) {
		counter := NewCounter()
		counter.AddItems("foo", "foo", "foo", "baz", "bar", "bar")
		assert.Equal(t, []CounterItem{
			{"foo", 3},
			{"bar", 2},
			{"baz", 1},
		}, counter.MostCommon(3))

		assert.Equal(t, []CounterItem{
			{"foo", 3},
			{"bar", 2},
			{"baz", 1},
		}, counter.MostCommon(3))

		assert.Equal(t, 3, counter.Get("foo"))
		assert.Equal(t, 2, counter.Get("bar"))
		assert.Equal(t, 1, counter.Get("baz"))
	})
}

func TestCounter_Get(t *testing.T) {
	t.Run("NonExistentItem", func(t *testing.T) {
		counter := NewCounter()
		assert.Equal(t, 0, counter.Get("foo"))
	})

	t.Run("ExistingItem", func(t *testing.T) {
		counter := NewCounter()
		counter.AddItems("foo", "bar", "foo")
		assert.Equal(t, 2, counter.Get("foo"))
		assert.Equal(t, 1, counter.Get("bar"))
	})
}
