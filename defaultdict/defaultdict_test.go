package defaultdict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultDict(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		d := New(IntDefault)
		d.Update("foo", func(value interface{}) interface{} {
			return value.(int) + 1
		})
		d.Update("foo", func(value interface{}) interface{} {
			return value.(int) + 1
		})

		assert.Equal(t, 2, d.Get("foo").(int))
	})

	t.Run("Int32", func(t *testing.T) {
		d := New(Int32Default)
		d.Update("foo", func(value interface{}) interface{} {
			return value.(int32) + 1
		})
		d.Update("foo", func(value interface{}) interface{} {
			return value.(int32) + 1
		})

		assert.Equal(t, int32(2), d.Get("foo").(int32))
	})

	t.Run("Int64", func(t *testing.T) {
		d := New(Int64Default)
		d.Update("foo", func(value interface{}) interface{} {
			return value.(int64) + 1
		})
		d.Update("foo", func(value interface{}) interface{} {
			return value.(int64) + 1
		})

		assert.Equal(t, int64(2), d.Get("foo").(int64))
	})

	t.Run("Float32", func(t *testing.T) {
		d := New(Float32Default)
		d.Update("foo", func(value interface{}) interface{} {
			return value.(float32) + 1
		})
		d.Update("foo", func(value interface{}) interface{} {
			return value.(float32) + 1
		})

		assert.Equal(t, float32(2), d.Get("foo").(float32))
	})

	t.Run("Float64", func(t *testing.T) {
		d := New(Float64Default)
		d.Update("foo", func(value interface{}) interface{} {
			return value.(float64) + 1
		})
		d.Update("foo", func(value interface{}) interface{} {
			return value.(float64) + 1
		})

		assert.Equal(t, float64(2), d.Get("foo").(float64))
	})

	t.Run("SetUpdateGet", func(t *testing.T) {
		d := New(IntDefault)
		d.Set("foo", 100)
		d.Set("bar", 10)

		assert.Equal(t, 100, d.Get("foo"))
		assert.Equal(t, 10, d.Get("bar"))

		d.Update("bar", func(value interface{}) interface{} {
			return 900
		})

		assert.Equal(t, 900, d.Get("bar"))
	})
}
