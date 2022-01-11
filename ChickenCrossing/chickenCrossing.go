package main

import (
	"fmt"
	"testing"
)

func main() {
	TestViewState(&testing.T{})
}
func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn mann ---\\__/ _________________/---]"
	state := ViewState(wanted)
	fmt.Println(state == wanted)
	// Checks if state is wanted
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	} else {
		fmt.Println(state)
	}
}

func MoveBoat(currentState string) string {
	// moves boat one character to the right
	oceanStart := len("[kylling rev korn mann ---")
	oceanEnd := len("[kylling rev korn mann ---\\__/ _________________")
	boat := currentState[26:30]
	// can use this for dynamic updating of boat position https://stackoverflow.com/questions/20827332/how-to-find-a-character-index-in-golang

	movedState := currentState[:oceanStart+1] + string("_") + boat + currentState[30:oceanEnd-1] + currentState[48:]
	fmt.Println("here", movedState)

	return movedState
}

func ViewState(currentState string) string {
	// returns new state for viewing
	newState := MoveBoat(currentState)
	return newState
}
