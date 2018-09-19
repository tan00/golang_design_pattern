package pattern

import (
	"fmt"
	"sync"
)

type ISingleton interface {
	SaySomething()
}

type Singleton struct {
}

func (Singleton) SaySomething() {
	fmt.Println("Singleton.SaySomething()")
}

var singletonInstance Singleton
var once sync.Once

func NewSingletonInstance() Singleton {
	once.Do(func() {
		singletonInstance = Singleton{}
	})

	return singletonInstance
}
