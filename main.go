package main

import (
	"fmt"

	"github.com/tizz98/go-playground/collections"
)

func main() {
	d := collections.NewOrderedDict()
	d.Set("foo", "bar")
	d.Set("baz", 123)

	for v := range d.Iterate() {
		fmt.Println(v)
	}
}

