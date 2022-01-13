package main

import (
	"fmt"
	"testing"
)

func main() {
	TestViewState(&testing.T{})
}
func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn mann ---\\\\__/ _________________/---]" // weird thing where double \\ is converted to single \
	state := ViewState(wanted)

	fmt.Println(state == wanted)
	// Checks if state is wanted
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	} else {
		fmt.Println(state)
	}
}

// function that moves boat to other shore
/**
Example of output
0 [kylling rev korn mann ---\_\__/ ________________/---]
1 [kylling rev korn mann ---\___\__/ _______________/---]
2 [kylling rev korn mann ---\_____\__/ ______________/---]
3 [kylling rev korn mann ---\_______\__/ _____________/---]
**/
func MoveBoat(currentState string) string {
	oceanStart := len("[kylling rev korn mann ---") + 1 // needed to add one more because character causes break statement
	oceanEnd := len("[kylling rev korn mann ---\\__/ _________________")

	// add functionality that allows you to select (replace thing with ----) "kylling", "rev" or something else to transport

	// displays start state
	fmt.Println(currentState)

	// max number to move boat is 18
	count := 18
	for i := 0; i < count; i++ {
		// updates the object by incrementing their positions by one as one "_" is added for each iteration
		// if ocean between boat is equal to shore start, there is no ocean
		if 30+i == oceanEnd-1 {
			startShore := currentState[:oceanStart+i]
			boat := currentState[27+i : 30+i+1]
			endShore := currentState[48+1:]

			movedState := startShore + string("_") + boat + endShore
			currentState = movedState
			fmt.Println(currentState)
		} else {
			startShore := currentState[:oceanStart+i]
			boat := currentState[27+i : 30+i]
			ocean := currentState[30+i : oceanEnd-1]
			endShore := currentState[48:]

			movedState := startShore + string("_") + boat + ocean + endShore
			currentState = movedState
			fmt.Println(currentState)
		}

	}
	// returns final state boat is in
	endState := currentState
	return endState
}

func ViewState(startState string) string {
	// returns new state for viewing
	newState := MoveBoat(startState)
	return newState
}
