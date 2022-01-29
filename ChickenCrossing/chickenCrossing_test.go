package main

import (
	"fmt"
	"testing"
)

func TestEndState(t *testing.T) {
	wantedState := RiverState{
		west: West{0: "dls", 1: "", 2: "", 3: ""},
		boat: Boat{0: "", 1: ""},
		east: East{0: "kylling", 1: "rev", 2: "mann", 3: "korn"},
	}

	currentState := &riverState

	// checks if values in states match
	for i := 0; i < len(currentState.west); i++ {
		if currentState.west[i] != wantedState.west[i] {
			t.Fatalf("Failed at index %d", i)
		}
	}
	for i := 0; i < len(currentState.boat); i++ {
		if currentState.boat[i] != wantedState.boat[i] {
			t.Fatalf("Failed at index %d", i)
		}
	}
	for i := 0; i < len(currentState.east); i++ {
		if currentState.east[i] != wantedState.east[i] {
			t.Fatalf("Failed at index %d", i)
		}
	}
}

func TestViewState(t *testing.T) {

	currentState := &riverState

	wantedGraphic := "[_kylling_rev_korn____W~~~~~~~~~~~~\\_____/mann|__/~~~~~~~~~~~~~~~E_______]"
	currenGraphic := ViewState(currentState)

	// Checks if state is wanted
	if currenGraphic != wantedGraphic {
		t.Errorf("Feil, fikk %q, ønsket %q.", currenGraphic, wantedGraphic)
	} else {
		fmt.Println(currenGraphic)
	}
}
