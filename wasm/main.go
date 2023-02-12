package main

import (
    "fmt"
	"syscall/js"
	//"taxi"
)

var document = js.Global().Get("document")

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func resetTaxiWorld(this js.Value, args []js.Value) interface{} {
	fmt.Println(taxi.ResetTaxiWorld())
	return nil
}


func sayHello(this js.Value, args []js.Value) any {
	inputValue := getElementByID("inputField").Get("value").String()
	outputValue := "Hello " + inputValue + "!"
	getElementByID("outputInfo").Set("innerHTML", outputValue)
	return nil
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
	js.Global().Set("sayHello", js.FuncOf(sayHello))
	//js.Global().Set("showTaxiWorld", js.FuncOf(showTaxiWorld))
	//js.Global().Set("updateTaxiLocation", js.FuncOf(updateTaxiLocation))
	//js.Global().Set("checkPassengerLocation", js.FuncOf(checkPassengerLocation))
	//js.Global().Set("getPassengerLocation", js.FuncOf(getPassengerLocation))

	<-done

}
