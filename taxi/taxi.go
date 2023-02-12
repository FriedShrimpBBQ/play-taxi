package taxi

import (
	"fmt"
	"strings"
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
		coordinates := UpdateTaxiCoordinates(taxiX, taxiY, act)
		return SetTaxiLocation(ResetTaxiWorld(), coordinates[0], coordinates[1])
	}
	return taxiWorld
}

func UpdateTaxiCoordinates(taxiX int, taxiY int, act int) []int {

	if act == 0 {
		taxiY = min(taxiY+1, 4)

	} else if act == 1 {
		taxiY = max(taxiY-1, 0)
	} else if act == 2 {
		taxiX = max(taxiX+1, 4)

	} else {
		taxiX = min(taxiX-1, 0)
	}

	return []int{taxiX, taxiY}
}

func ShowTaxiWorld(taxiWorld [][]string) {
	fmt.Print(TaxiWorldStringify(taxiWorld))
}

func TaxiWorldStringify(taxiWorld [][]string) string {
	taxiWorldFlatten := []string{}
	for r, _ := range taxiWorld {
		taxiWorldFlatten = append(taxiWorldFlatten, strings.Join(taxiWorld[r], ""))
	}
	return strings.Join(taxiWorldFlatten, "\n")
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
