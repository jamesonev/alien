package main

import (
	"fmt"
	"strings"
)

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
func printAtlas(atlas map[string]*city) {
	for k := range atlas {
		printCity(atlas[k])
	}
}

// addDirection makes the assumption that if Foo.north is Bar, then Bar.south should be Foo
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
	//make the string lowercase and only compare against the first letter to handle multiple
	//ways of expressing the same direction: "north", "North", "N", etc
	direction = strings.ToLower(direction)
	if strings.HasPrefix(direction, "n") {
		setNorth(sourceCity, destCity)
	} else if strings.HasPrefix(direction, "s") {
		setSouth(sourceCity, destCity)
	} else if strings.HasPrefix(direction, "e") {
		setEast(sourceCity, destCity)
	} else if strings.HasPrefix(direction, "w") {
		setWest(sourceCity, destCity)
	}

}

func main() {
	var atlas = make(map[string]*city)
	addDirection(atlas, "Foo", "Bar", "north")
	addDirection(atlas, "Boo", "Bar", "east")
	addDirection(atlas, "Foo", "Boo", "west")
	printAtlas(atlas)
}
