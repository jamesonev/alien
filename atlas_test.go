package alien

import "testing"

func TestRemoveLinks(t *testing.T) {
	atlas := parseFile("input.txt")
	removeLinks(atlas["Bar"])
	neighbor := getNeighbor(atlas["Bar"])
	if neighbor != nil {
		t.Error("neighbor found after call to removeLinks")
	}
}

func TestAddDirection(t *testing.T) {
	atlas := parseFile("input.txt")
	addDirection(atlas, "Bar", "Baz", "north")
	if atlas["Baz"].s.name != "Bar" {
		t.Error("Back link not created")
	}
}
