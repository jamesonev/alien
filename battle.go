package alien

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type warzone struct {
	city       *City
	occupiedBy int
}

func newWarzone(city *City) *warzone {
	new := warzone{city: city, occupiedBy: -1}
	return &new
}

// this handles an alien arriving in a city and returns whether the city was destroyed
func alienArrivesInCity(index, alien int, battle []*warzone, atlas map[string]*City) bool {
	wasDestroyed := false
	index = index % len(battle)
	current := battle[index]
	//if there's no one in the city, put this alien there
	if current != nil && current.city != nil {
		// the city is empty, the alien moves in
		if current.occupiedBy == -1 {
			current.occupiedBy = alien
		} else if current.occupiedBy != alien { //an alien can't arrive in the same city as its already in
			// an alien is already here so they have a fight
			fmt.Print("Oh no! ", current.city.name, " was destroyed by alien ")
			fmt.Println(current.occupiedBy, "and alien", alien, "!")
			//remove links and backlinks
			removeLinks(current.city)
			//remove from our atlas
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

func checkForAliens(battle []*warzone) {
	fmt.Println("checking for hiding aliens")
	for _, zone := range battle {
		if zone != nil {
			if zone.occupiedBy != -1 {
				fmt.Println("Alien", zone.occupiedBy, "is in", zone.city.name)
			}
		}
	}
}

// this function processes the command line args and returns the number of aliens specified,
// and the string format file name if one was specified. defaults to 'input.txt'
func processArgs() (int, string) {
	//get command line args
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
	return numAliens, fileName
}

// Attack simulates aliens arriving in various cities as described in task.md
func Attack() {
	numAliens, fileName := processArgs()
	atlas := parseFile(fileName)
	// printAtlas(atlas)
	fmt.Println("here's the earth did")
	// Here's where the war starts
	// 'battle' is a slice which holds pointers to warzones. storing warzones in a slice makes randomly assigning
	// aliens to cities easy
	battleSize := len(atlas)
	battle := make([]*warzone, battleSize)
	// lookup gives us a way to map city pointers, which return a name, back to a position in the battle
	// that way, when an alien is going to visit a neighbor city, we can move them there in constant time
	lookup := make(map[string]int)
	i := 0
	for k := range atlas {
		// i is the index, where k is a string
		battle[i] = newWarzone(atlas[k])
		lookup[atlas[k].name] = i
		i++
	}
	aliensAlive := numAliens
	//drop the aliens into random cities
	for i := 0; i < numAliens; i++ {
		// we need to keep track of how many aliens are alive so we know if we can exit early
		if alienArrivesInCity(rand.Int(), i, battle, atlas) == true {
			aliensAlive -= 2
		}
	}
	// from the spec, simulate 10000 moves per alien
	for i := 0; i < 10000; i++ {
		if aliensAlive == 0 {
			fmt.Println("All the aliens are dead!")
			printAtlas(atlas)
			return
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
						aliensAlive -= 2
					}
					//we moved the alien, so we need to flag the city is open
					current.occupiedBy = -1
				}

			}
		}
	}
	checkForAliens(battle)
	printAtlas(atlas)
}
