package pattern

import "fmt"

//抽象工厂
type AbstractFactory interface {
	Producepen(pentype string) pen    //生产笔
	Produceball(balltype string) ball //生产球
}

type FactorProducer struct {
}

func (*FactorProducer) GetFactory(name string) AbstractFactory {
	switch name {
	case "FactoryA":
		return new(FactoryA)
	case "FactoryB":
		return new(FactoryB)
	default:
		fmt.Println("not support Factory")
	}
	return nil
}

//球类
type ball interface {
	Play()
}

//篮球
type basketBall struct {
}

//篮球
type footBall struct {
}

func (*basketBall) Play() {
	fmt.Println("basketBall.Play")
}
func (*footBall) Play() {
	fmt.Println("footBall.Play")
}

//FactoryA  可以生产 pencil basketball
type FactoryA struct {
}

func (*FactoryA) Producepen(pentype string) pen {
	switch pentype {
	case "pencil":
		return new(pencil)
	default:
		fmt.Println("not support product")
		return nil
	}
}
func (*FactoryA) Produceball(balltype string) ball {
	switch balltype {
	case "basketball":
		return new(basketBall)

	default:
		fmt.Println("not support product")
		return nil
	}
}

//FactoryB  可以生产 brush football
type FactoryB struct {
}

func (*FactoryB) Producepen(pentype string) pen {
	switch pentype {
	case "brush":
		return new(brushPen)
	default:
		fmt.Println("not support product")
		return nil
	}
}
func (*FactoryB) Produceball(balltype string) ball {
	switch balltype {
	case "football":
		return new(footBall)
	default:
		fmt.Println("not support product")
		return nil
	}
}
