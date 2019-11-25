package alien

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

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

func Attack() {
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
