package main

import (
	"flag"
	"fmt"

	"./pattern"
)

// var (
// 	patternName = flag.String("", "", "Which  Pattern  to Run")
// )

func adapter() {
	var battery pattern.RechargeableBattery
	battery = pattern.AdapterNonToYes{pattern.NonRechargeableA{}}
	battery.Use()
	battery.Charge()

	battery = pattern.NonRechargeableB{}
	battery.Use()
	battery.Charge()
}

func factory() {
	fac := new(pattern.PenFactory)
	pen := fac.Produce("pencil")
	pen.Write()

	pen2 := fac.Produce("brush")
	pen2.Write()

}

func abstractFactory() {
	facA := new(pattern.FactoryA)
	facA.Producepen("pencil").Write()
	facA.Produceball("basketball").Play()

	facB := new(pattern.FactoryB)
	facB.Producepen("brush").Write()
	facB.Produceball("football").Play()
}

func main() {

	flag.Parse()
	patternName := flag.Arg(0)

	fmt.Println("pattern name = " + patternName)

	switch patternName {

	case "Factory":
		factory()
	case "Adapter":
		adapter()
	case "AbstractFactory":
		abstractFactory()
	}

}
