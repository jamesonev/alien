package alien

import (
	"bufio"
	"fmt"
	"math/rand"
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

//this function randomly picks a direction and then returns the first valid city pointer it finds
func getNeighbor(city *City) *City {
	index := rand.Intn(4)
	for i := 0; i < 4; i++ {
		switch index + i%4 {
		case 0:
			if city.n != nil {
				return city.n
			}
		case 1:
			if city.s != nil {
				return city.s
			}
		case 2:
			if city.e != nil {
				return city.e
			}
		case 3:
			if city.w != nil {
				return city.w
			}
		}
	}
	return nil
}

// these are helpers for setDirection(). Because ensuring that cities have correct complementary pointers
// is so important, I built all these instead of setting by hand
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

// this just prints cities in the specified format
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
	// we begin by checking if the src and dest cities exist, and create them if they don't
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

// this function sets all direction pointers to nil. It also follows all those pointers
// to set the backlinks to be nil
func removeLinks(city *City) {
	if city.n != nil {
		north := city.n
		north.s = nil
		city.n = nil
	}
	if city.s != nil {
		south := city.s
		south.n = nil
		city.s = nil
	}
	if city.e != nil {
		east := city.e
		east.w = nil
		city.e = nil
	}
	if city.w != nil {
		west := city.w
		west.e = nil
		city.w = nil
	}
}

// parsefile takes the input filename and makes a map of the different cities
//
func parseFile(fileName string) map[string]*City {
	var atlas = make(map[string]*City)
	var err error
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	success := scanner.Scan()
	for success != false {
		line := scanner.Text()
		fields := strings.Fields(line)
		currentCity := fields[0]
		//check if we have a direction to add
		if len(fields) > 1 {
			directions := fields[1:]
			for _, direction := range directions {
				pair := strings.Split(direction, "=")
				direction = pair[0]
				city := pair[1]
				addDirection(atlas, currentCity, city, direction)
			}
		} else {
			//TODO: add handling for the case where a city doesn't have any directions
			fmt.Println("Error: city", currentCity, "has no cities associated with it")
		}
		success = scanner.Scan()
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	return atlas
}
