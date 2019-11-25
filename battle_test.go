package alien

import (
	"testing"
)

func TestAlienArrives(t *testing.T) {
	atlas := parseFile("input.txt")
	battle := make([]*warzone, len(atlas))
	i := 0
	for k := range atlas {
		// i is the index, where k is a string
		battle[i] = newWarzone(atlas[k])
		i++
	}
	if alienArrivesInCity(0, 0, battle, atlas) {
		t.Errorf("city was destroyed by first alien arrival")
	}
	if !alienArrivesInCity(0, 1, battle, atlas) {
		t.Errorf("city was not destroyed")
	}
}
