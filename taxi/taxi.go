package taxi

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func ResetTaxiWorld() [][]string {
	return [][]string{
		{"+", "-", "-", "-", "-", "-", "-", "-", "-", "-", "+"},
		{"|", "R", ":", " ", "|", " ", ":", " ", ":", "G", "|"},
		{"|", " ", ":", " ", "|", " ", ":", " ", ":", " ", "|"},
		{"|", " ", ":", " ", ":", " ", ":", " ", ":", " ", "|"},
		{"|", " ", "|", " ", ":", " ", "|", " ", ":", " ", "|"},
		{"|", "Y", "|", " ", ":", " ", "|", "B", ":", " ", "|"},
		{"+", "-", "-", "-", "-", "-", "-", "-", "-", "-", "+"},
	}
}

func SetTaxiLocation(taxiWorld [][]string, taxiX int, taxiY int) [][]string {
	if taxiX >= 0 && taxiX <= 4 {
		// we update the taxi location with "T" using offset if valid
		taxiWorld[taxiY+1][taxiX*2+1] = "T"
	}
	return taxiWorld

}

func UpdateTaxiLocation(taxiWorld [][]string, taxiX int, taxiY int, act int) [][]string {
	if taxiWorld[taxiY+1][taxiX*2+1] == "T" {
		if act == 0 {
			taxiY = min(taxiY+1, 4)
			return SetTaxiLocation(ResetTaxiWorld(), taxiX, taxiY)
		} else if act == 1 {
			taxiY = max(taxiY-1, 0)
			return SetTaxiLocation(ResetTaxiWorld(), taxiX, taxiY)
		} else if act == 2 {
			taxiX = max(taxiX+1, 4)
			return SetTaxiLocation(ResetTaxiWorld(), taxiX, taxiY)
		} else {
			taxiX = min(taxiX-1, 0)
			return SetTaxiLocation(ResetTaxiWorld(), taxiX, taxiY)
		}
	}
	return taxiWorld
}

func ShowTaxiWorld(taxiWorld [][]string) {
	for r, _ := range taxiWorld {

		// Then we iterate over the items of each row:
		for _, colValue := range taxiWorld[r] {

			// See string formatting docs at
			// https://gobyexample.com/string-formatting
			fmt.Print(colValue)
		}
		fmt.Print("\n")
	}
}

func CheckPassengerLocation(taxiWorld [][]string, passenger int, goal int) int {
	// checks where the passenger is on the taxiWorld and verify if Goal has been reached or not
	passengerMapping := map[int]string{0: "R", 1: "G", 2: "Y", 3: "B", 4: "T"}
	goalMapping := map[int]string{0: "R", 1: "G", 2: "Y", 3: "B"}
	if passengerMapping[passenger] == "R" && taxiWorld[1][1] == "T" {
		passenger = 4
	} else if passengerMapping[passenger] == "Y" && taxiWorld[5][1] == "T" {
		passenger = 4
	} else if passengerMapping[passenger] == "G" && taxiWorld[1][9] == "T" {
		passenger = 4
	} else if passengerMapping[passenger] == "B" && taxiWorld[5][7] == "T" {
		passenger = 4
	}

	// now check if we reached the goal
	if passenger == 4 && goalMapping[goal] == "R" && taxiWorld[1][1] == "T" {
		passenger = 5
	} else if passenger == 4 && goalMapping[goal] == "Y" && taxiWorld[5][1] == "T" {
		passenger = 5
	} else if passenger == 4 && goalMapping[goal] == "G" && taxiWorld[1][9] == "T" {
		passenger = 5
	} else if passenger == 4 && goalMapping[goal] == "B" && taxiWorld[5][7] == "T" {
		passenger = 5
	}
	return passenger
}

func GetPassengerLocation(passenger int) string {
	return map[int]string{0: "R", 1: "G", 2: "Y", 3: "B", 4: "T"}[passenger]
}
