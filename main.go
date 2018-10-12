package main

import (
	"fmt"

	"github.com/tan00/golang_design_pattern/pattern"
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

	facproducor := new(pattern.FactorProducer)

	facA := facproducor.GetFactory("FactoryA")
	facA.Producepen("pencil").Write()
	//facA.Producepen("pen not support").Write()
	facA.Produceball("basketball").Play()

	facB := facproducor.GetFactory("FactoryB")
	facB.Producepen("brush").Write()
	facB.Produceball("football").Play()
}
func singleton() {
	instance := pattern.NewSingletonInstance()
	instance.SaySomething()

	instance2 := pattern.NewSingletonInstance()

	if instance != instance2 {
		fmt.Println("singleton failed")
	}
}

func builder() {
	pattern.BuilderClient()
}

const (
	idFactory = iota + 1
	idAbsFactory
	idSingleton
	idPrototype
	idBuilder
)

func main() {
	var (
		//err       error
		idPattern int
		menu      string
	)
	menu = `
	*1. factory 
	*2. abstract factroy
	*3. singleton
	*4. prototype  原型模式: 对象的复制
	*5. builder 
	`

	fmt.Printf("select pattern: \n %s \n ", menu)
	fmt.Scanf("%d", &idPattern)

	switch idPattern {
	case idFactory:
		factory()
	case idAbsFactory:
		abstractFactory()
	case idSingleton:
		singleton()
	case idBuilder:
		builder()
	default:
		fmt.Println("undefined function")
	}

}
