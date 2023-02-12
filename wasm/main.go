package main

import (
	"fmt"
	"syscall/js"
	"taxi"
)

var document = js.Global().Get("document")

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func resetTaxiWorld(this js.Value, args []js.Value) any {
	world := taxi.TaxiWorldStringify(taxi.ResetTaxiWorld())
	fmt.Println(world)
	return world
}

func sayHello(this js.Value, args []js.Value) any {
	inputValue := getElementByID("inputField").Get("value").String()
	outputValue := "Hello " + inputValue + "!"
	getElementByID("outputInfo").Set("innerHTML", outputValue)
	return nil
}

func updateTaxiCoordinates(this js.Value, args []js.Value) any {

	taxiX := args[0].Get("taxiX").Int()
	taxiY := args[0].Get("taxiY").Int()
	act := args[0].Get("act").Int()

	fmt.Println(taxiX, taxiY, act)

	if act >= 0 {
		coordinates := taxi.UpdateTaxiCoordinates(taxiX, taxiY, act)
		fmt.Println("from golang", taxiX, taxiY, coordinates)
		return js.ValueOf(map[string]interface{}{
			"taxiX": coordinates[0],
			"taxiY": coordinates[1],
		})

	} else {
		return js.ValueOf(map[string]interface{}{
			"taxiX": taxiX,
			"taxiY": taxiY,
		})
	}
}

func taxiWorldStringify(this js.Value, args []js.Value) any {

	taxiX := args[0].Get("taxiX").Int()
	taxiY := args[0].Get("taxiY").Int()

	world := taxi.SetTaxiLocation(taxi.ResetTaxiWorld(), taxiX, taxiY)
	return taxi.TaxiWorldStringify(world)
}

//func showTaxiWorld(this js.Value, args []js.Value) interface{} {
//	return taxi.ShowTaxiWorld()
//}

//func updateTaxiLocation(this js.Value, args []js.Value) interface{} {
//	taxi.UpdateTaxiLocation
//}

//func checkPassengerLocation(this js.Value, args []js.Value) interface{} {
//	taxi.CheckPassengerLocation
//}

//func getPassengerLocation(this js.Value, args []js.Value) interface{} {
//	taxi.GetPassengerLocation
//}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("resetTaxiWorld", js.FuncOf(resetTaxiWorld))
	js.Global().Set("updateTaxiCoordinates", js.FuncOf(updateTaxiCoordinates))
	js.Global().Set("taxiWorldStringify", js.FuncOf(taxiWorldStringify))
	js.Global().Set("sayHello", js.FuncOf(sayHello))
	//js.Global().Set("showTaxiWorld", js.FuncOf(showTaxiWorld))
	//js.Global().Set("updateTaxiLocation", js.FuncOf(updateTaxiLocation))
	//js.Global().Set("checkPassengerLocation", js.FuncOf(checkPassengerLocation))
	//js.Global().Set("getPassengerLocation", js.FuncOf(getPassengerLocation))

	<-done

}
