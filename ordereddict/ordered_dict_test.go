package ordereddict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedDict(t *testing.T) {
	d := New()
	d.Set("foo", "bar")
	d.Set("test", 123)

	assert.Equal(t, d.Get("foo"), "bar")
	assert.Equal(t, d.Get("test"), 123)

	assert.True(t, d.Remove("test"))
	assert.False(t, d.Remove("test"))

	d.Set("test", 123)
	d.Set("baz", "baz")
	d.Set("test", 456)
	assert.Equal(t, d.Get("test"), 456)

	// because we call .Set("test", ...) again, it should be moved
	// to the back of the list
	t.Run("Iterate", func(t *testing.T) {
		i := 0
		expected := []interface{}{"bar", "baz", 456}

		for v := range d.Iterate() {
			assert.Equal(t, expected[i], v)
			i++
		}
	})
}

