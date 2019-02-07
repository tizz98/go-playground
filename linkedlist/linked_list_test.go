package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("Iteration", func(t *testing.T) {
		l := New()
		l.Append("foo")
		l.Append("Bar")

		i := 0
		expected := []string{"foo", "Bar"}

		for n := range l.Iterate() {
			assert.Equal(t, n.value.(string), expected[i])
			i++
		}
	})

	t.Run("Removal", func(t *testing.T) {
		l := New()
		l.Append("foo")
		bar := l.Append("Bar")
		l.Append("baz")
		l.Append("abc")

		assert.True(t, l.Remove(bar))

		i := 0
		expected := []string{"foo", "baz", "abc"}

		for n := range l.Iterate() {
			assert.Equal(t, n.value.(string), expected[i])
			i++
		}
	})
}

