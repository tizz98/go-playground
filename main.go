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

	c := collections.NewCounter()
	c.AddItems("foo", "foo", "bar", "baz")

	for i, item := range c.MostCommon(5) {
		fmt.Printf("%d: %#v seen %d time(s)\n", i+1, item.Value, item.Count)
	}
}
