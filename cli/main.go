package main

import (
	"flag"
	"fmt"
	"taxi"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

var taxiWorld = taxi.ResetTaxiWorld()

func main() {
	var taxiX, taxiY, act, passenger, goal int
	flag.IntVar(&taxiX, "taxi-x", 0, "Taxi X coordinate")
	flag.IntVar(&taxiY, "taxi-y", 0, "Taxi X coordinate")
	flag.IntVar(&act, "act", 0, "Taxi Move Direction")
	flag.IntVar(&passenger, "passenger", 0, "Passenger Location")
	flag.IntVar(&goal, "goal", 0, "Goal Location")
	flag.Parse()

	if isFlagPassed("taxi-x") && isFlagPassed("taxi-y") {
		taxiWorld = taxi.SetTaxiLocation(taxiWorld, taxiX, taxiY)
		fmt.Println("")
	} else {
		fmt.Printf("Flags are not set\n")
	}

	taxi.ShowTaxiWorld(taxiWorld)

	if isFlagPassed("act") {
		fmt.Println("\n~Taxi Move~")
		taxiWorld = taxi.UpdateTaxiLocation(taxiWorld, taxiX, taxiY, act)
		taxi.ShowTaxiWorld(taxiWorld)
	}

	if isFlagPassed("passenger") && isFlagPassed("goal") {
		passenger := taxi.CheckPassengerLocation(taxiWorld, passenger, goal)
		fmt.Println("Passegner location:", map[int]string{0: "R", 1: "G", 2: "Y", 3: "B", 4: "T"}[passenger])
	}

}
