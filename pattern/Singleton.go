package pattern

import (
	"fmt"
	"sync"
)

type Singleton interface {
	SaySomething()
}

type singleton struct {
}

func (singleton) SaySomething() {
	fmt.Println("Singleton")
}

var singletonInstance Singleton
var once sync.Once

func NewSingletonInstance() Singleton {
	once.Do(func() {
		singletonInstance = &singleton{}
	})

	return singletonInstance
}
