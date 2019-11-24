package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type City struct {
	name string
	n    *City
	s    *City
	e    *City
	w    *City
}

func newCity(name string) *City {
	new := City{name: name}
	return &new
}

//these helpers will keep our atlas in sync but will only be available through SetDirection()
func setNorth(src, dest *City) {
	src.n = dest
	if dest != nil {
		dest.s = src
	}
}
func setSouth(src, dest *City) {
	src.s = dest
	if dest != nil {
		dest.n = src
	}
}
func setEast(src, dest *City) {
	src.e = dest
	if dest != nil {
		dest.w = src
	}
}
func setWest(src, dest *City) {
	src.w = dest
	if dest != nil {
		dest.e = src
	}
}

func printCity(c *City) {
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
func printAtlas(atlas map[string]*City) {
	for k := range atlas {
		printCity(atlas[k])
	}
}

// addDirection makes the assumption that if Foo.north is Bar, then Bar.south should be Foo
func addDirection(atlas map[string]*City, src, dest, direction string) {
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

	direction = strings.ToLower(strings.TrimSpace(direction))
	if strings.HasPrefix(direction, "n") {
		setNorth(sourceCity, destCity)
	} else if strings.HasPrefix(direction, "s") {
		setSouth(sourceCity, destCity)
	} else if strings.HasPrefix(direction, "e") {
		setEast(sourceCity, destCity)
	} else if strings.HasPrefix(direction, "w") {
		setWest(sourceCity, destCity)
	} else {
		panic(direction)
	}

}

func parseFile(fileName string) map[string]*City {
	var atlas = make(map[string]*City)
	var err error
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	success := scanner.Scan()
	for success != false {
		line := scanner.Text()
		fields := strings.Fields(line)
		currentCity := fields[0]
		directions := fields[1:]
		for _, direction := range directions {
			pair := strings.Split(direction, "=")
			direction = pair[0]
			city := pair[1]
			addDirection(atlas, currentCity, city, direction)
		}
		success = scanner.Scan()
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	return atlas
}

func main() {
	argv := os.Args
	if len(argv) != 3 {
		panic("must provide 2 command line args, numAliens and inputfile")
	}
	// numAliens, err := strconv.Atoi(argv[1])
	// if err != nil {
	// 	panic(err)
	// }
	fileName := strings.TrimSpace(argv[2])
	atlas := parseFile(fileName)
	printAtlas(atlas)
}
