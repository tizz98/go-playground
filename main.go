package main

import (
	"fmt"

	"github.com/tizz98/go-playground/counter"
	"github.com/tizz98/go-playground/defaultdict"
	"github.com/tizz98/go-playground/ordereddict"
)

func main() {
	d := ordereddict.New()
	d.Set("foo", "bar")
	d.Set("baz", 123)

	for v := range d.Iterate() {
		fmt.Println(v)
	}

	c := counter.New()
	c.AddItems("foo", "foo", "bar", "baz")

	for i, item := range c.MostCommon(5) {
		fmt.Printf("%d: %#v seen %d time(s)\n", i+1, item.Value, item.Count)
	}

	dict := defaultdict.New(defaultdict.IntDefault)
	fmt.Printf("foo default: %d\n", dict.Get("foo"))
}
