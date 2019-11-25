package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

// this function sets all direction pointers to nil. It also follows all those pointers
// to set the backlinks to be nil
func destroyCity(city *City) {
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
				fmt.Println(currentCity)
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

type Warzone struct {
	city       *City
	occupiedBy int
}

func newWarzone(city *City) *Warzone {
	new := Warzone{city: city, occupiedBy: -1}
	return &new
}
func alienArrivesInCity(index, alien int, battle []*Warzone, atlas map[string]*City) bool {
	wasDestroyed := false
	index = index % len(battle)
	current := battle[index]
	//if there's no one in the city, put this alien there
	if current != nil && current.city != nil {
		if current.occupiedBy == -1 {
			current.occupiedBy = alien
		} else if current.occupiedBy != alien { //an alien can't arrive in the same city it is in
			// an alien is already here so they have a fight
			fmt.Print("Oh no! ", current.city.name, " was destroyed by alien ")
			fmt.Println(current.occupiedBy, "and alien", alien, "!")
			//remove links and backlinks
			destroyCity(current.city)
			//remove form our atlas
			delete(atlas, current.city.name)
			//remove other references to mark for garbage collection
			current.occupiedBy = -1
			current.city = nil
			battle[index] = nil
			wasDestroyed = true
		}
	}
	return wasDestroyed
}
func main() {
	argv := os.Args
	fileName := "input.txt"
	if len(argv) < 2 {
		panic("must provide 1 command line arg, numAliens. input file optional")
	}
	if len(argv) == 3 {
		fileName = strings.TrimSpace(argv[2])
	}
	numAliens, err := strconv.Atoi(argv[1])
	if err != nil {
		panic(err)
	}
	atlas := parseFile(fileName)
	printAtlas(atlas)
	fmt.Println("here's the aftermath")
	// here's where the war starts
	//'battle' is a slice which holds pointers to warzones. initially, we
	battleSize := len(atlas)
	battle := make([]*Warzone, battleSize)
	lookup := make(map[string]int)
	i := 0
	for k := range atlas {
		battle[i] = newWarzone(atlas[k])
		lookup[atlas[k].name] = i
		i++
	}
	//drop the aliens into random cities
	for i := 0; i < numAliens; i++ {
		// we need to keep track of how many aliens are alive so we know if we can exit early
		if alienArrivesInCity(rand.Int(), i, battle, atlas) == true {
			numAliens -= 2
			battleSize--
		}
	}
	// from the spec, simulate 10000 moves per alien
	for i := 0; i < 10000; i++ {
		if numAliens == 0 {
			fmt.Println("All the aliens are dead!")
			break
		}
		for j := 0; j < battleSize; j++ {
			//iterate over the battle, and move each alien we find
			current := battle[j]
			//find an alien
			if current != nil && current.occupiedBy != -1 {
				alien := current.occupiedBy
				//get a city for it to move into
				neighbor := getNeighbor(current.city)
				//simulate it arriving in that city
				if neighbor != nil {
					if alienArrivesInCity(lookup[neighbor.name], alien, battle, atlas) {
						numAliens -= 2
						battleSize--
					}
				}

			}
		}
	}

	printAtlas(atlas)
}
