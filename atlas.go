package main

import "fmt"

type city struct {
	name string
	n    *city
	s    *city
	e    *city
	w    *city
}

func newCity(name string, north, south, east, west *city) *city {
	new := city{name: name, n: north, s: south, e: east, w: west}
	return &new
}

func printCity(c *city) {
	fmt.Print(c.name)
	if c.n != nil {
		fmt.Print(" north=", c.n.name)
	}
	if c.s != nil {
		fmt.Print(" south=", c.s.name)
	}
	if c.e != nil {
		fmt.Print(" east=", c.e.name)
	}
	if c.w != nil {
		fmt.Print(" west=", c.w.name)
	}
	fmt.Println()
}

func main() {
	var atlas = make(map[string]*city)
	atlas["Foo"] = newCity("Foo", nil, nil, nil, nil)
	atlas["Bar"] = newCity("Bar", nil, atlas["Foo"], nil, nil)
	printCity(atlas["Bar"])
}
