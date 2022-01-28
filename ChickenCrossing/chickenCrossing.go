package main

import (
	"fmt"
	"testing"
	"time"
)

type West []string
type Boat []string
type East []string

type RiverState struct {
	east []string
	boat []string
	west []string
}

var riverState = RiverState{
	west: West{0: "kylling", 1: "rev", 2: "korn", 3: "mann"},
	boat: Boat{0: "", 1: ""},
	east: East{0: "", 1: "", 2: "", 3: ""},
}

func main() {
	// moves from west
	MoveToBoat(&riverState, riverState.west[0], riverState.west[3])

	// to east
	UnloadEast(&riverState, 0, 3)

	// then loads with no cargo and captain. then moves to west
	MoveToBoat(&riverState, riverState.boat[0], riverState.east[3])

	// unloads no cargo to west side

	// then loads rev as cargo from west keeping captain
	MoveToBoat(&riverState, riverState.west[1], riverState.boat[1])

	// to east
	UnloadEast(&riverState, 1, 3) //

	// loads kylling as cargo and mann as captain from east
	MoveToBoat(&riverState, riverState.east[0], riverState.east[3])

	// to west

	// unloads kylling and mann, before loading korn as cargo and mann as captain from west
	UnloadWest(&riverState, 0, 3)

	MoveToBoat(&riverState, riverState.west[2], riverState.west[3])

	// then unloads on east side, before the mann captains with no cargo
	UnloadEast(&riverState, 2, 3)

	MoveToBoat(&riverState, riverState.east[0], riverState.east[3])

	// then loads chicken as cargo from west, while keeping captain on boat
	MoveToBoat(&riverState, riverState.west[0], riverState.boat[1])

	// goes to east

	// then unloads chicken and man on east side
	UnloadEast(&riverState, 0, 3)

}

func MoveToBoat(r *RiverState, cargo string, captain string) {
	if len(r.boat[0]) == 0 {
		r.boat[0] = cargo
		r.boat[1] = captain

		// empties land where cargo and captain was loaded from
		for i := 0; i < len(r.west); i++ {
			if r.west[i] == cargo || r.west[i] == captain {
				r.west[i] = ""
			} else if r.east[i] == cargo || r.east[i] == captain {
				r.east[i] = ""
			}
		}

		println("\nPut ", cargo, " in boat with ", captain)
		println("cargo  : ", r.boat[0])
		println("captain: ", r.boat[1])

		time.Sleep(1 * time.Second)
	} else {
		println("Boat is full")
	}
	// prints current state
	PrintState(r)
}

func UnloadEast(r *RiverState, place1 int, place2 int) {
	if len(r.boat[0]) != 0 || len(r.boat[1]) != 0 {
		// prepares for unloading
		cargo := r.boat[0]
		captain := r.boat[1]
		landingPlace1 := r.east[place1]
		landingPlace2 := r.east[place2]

		// swaps cargo and captain with land items
		r.boat[0] = landingPlace1
		r.boat[1] = landingPlace2
		r.east[place1] = cargo
		r.east[place2] = captain

		println("\nUnloaded ", cargo, " and ", captain, " from boat in East")
		println("cargo  : ", r.boat[0])
		println("captain: ", r.boat[1])

		time.Sleep(1 * time.Second)

	} else {
		println("Boat is missing cargo or captain")
	}
	// prints current state
	PrintState(r)
}

func UnloadWest(r *RiverState, place1 int, place2 int) {
	if len(r.boat[0]) != 0 {
		// prepares for unloading
		cargo := r.boat[0]
		captain := r.boat[1]
		// swaps cargo and captain with land items
		landingPlace1 := r.west[place1]
		landingPlace2 := r.west[place2]

		r.boat[0] = landingPlace1
		r.boat[1] = landingPlace2
		r.west[place1] = cargo
		r.west[place2] = captain

		println("\nUnloaded ", cargo, " and ", captain, " from boat in West")
		println("cargo  : ", r.boat[0])
		println("captain: ", r.boat[1])

		time.Sleep(1 * time.Second)
	} else {
		println("Boat is missing cargo or captain")
	}
	// prints current state
	PrintState(r)
}

func PrintState(r *RiverState) {
	// prints current state in graphic format
	fmt.Printf(
		"[_%s_%s_%s_%s___W~~~~~~~~~~~~\\__%s___/%s|__/~~~~~~~~~~~~~~~E___%s_%s_%s_%s_]\n\n",
		r.west[0], r.west[1], r.west[2], r.west[3],
		r.boat[0], r.boat[1],
		r.east[0], r.east[1], r.east[2], r.east[3],
	)
}

// backup print of state
func PrintStateTicket(r *RiverState) {
	println("----------------------------------\nWest side:")
	for i := 0; i < len(r.west); i++ {
		println(r.west[i])
	}
	println("\nBoat contains:")
	for i := 0; i < len(r.boat); i++ {
		println(r.boat[i])
	}

	println("\nEast side:")
	for i := 0; i < len(r.east); i++ {
		println(r.east[i])
	}
	println("----------------------------------")
}

//

//

//

//

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

// function that animates boat moving to other shore
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
	boatLeftState := "\\\\__/ _________________/"
	// add functionality that allows you to select (replace thing with ----) "kylling", "rev" or something else to transport

	// displays start state
	fmt.Println(currentState)
	if currentState[oceanStart:oceanEnd+1] == boatLeftState {
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
