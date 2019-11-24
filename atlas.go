package main

import "fmt"

type city struct {
	name string
	n    *city
	s    *city
	e    *city
	w    *city
}

func newCity(name string) *city {
	new := city{name: name}
	return &new
}

//these helpers will keep our atlas in sync but will only be available through SetDirection()
func setNorth(src, dest *city) {
	src.n = dest
	if dest != nil {
		dest.s = src
	}
}
func setSouth(src, dest *city) {
	src.s = dest
	if dest != nil {
		dest.n = src
	}
}
func setEast(src, dest *city) {
	src.e = dest
	if dest != nil {
		dest.w = src
	}
}
func setWest(src, dest *city) {
	src.w = dest
	if dest != nil {
		dest.e = src
	}
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

func addDirection(atlas map[string]*city, src, dest, direction string) {
	sourceCity, exists := atlas[src]
	if !exists {
		atlas[src] = newCity(src)
		sourceCity = atlas[src]
	}
	destCity, exists := atlas[dest]
	if !exists {
		atlas[dest] = newCity(dest)
		destCity = atlas[dest]
	}
	if direction == "north" {
		setNorth(sourceCity, destCity)
	} else if direction == "south" {
		setSouth(sourceCity, destCity)
	} else if direction == "east" {
		setEast(sourceCity, destCity)
	} else if direction == "west" {
		setWest(sourceCity, destCity)
	}

}

func main() {
	var atlas = make(map[string]*city)
	// atlas["Foo"] = newCity("Foo")
	// atlas["Bar"] = newCity("Bar")
	addDirection(atlas, "Foo", "Bar", "north")
	printCity(atlas["Foo"])
}
