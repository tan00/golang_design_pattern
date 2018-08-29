package pattern

import "fmt"

//抽象工厂
type AbstractFactory interface {
	Producepen(pentype string) pen    //生产笔
	Produceball(balltype string) ball //生产球
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

func (this *basketBall) Play() {
	fmt.Println("basketBall.Play")
}
func (this *footBall) Play() {
	fmt.Println("footBall.Play")
}

//工厂A
type FactoryA struct {
}

func (this *FactoryA) Producepen(pentype string) pen {
	switch pentype {
	case "pencil":
		return new(pencil)
	default:
		return nil
	}
}
func (*FactoryA) Produceball(balltype string) ball {
	switch balltype {
	case "basketball":
		return new(basketBall)

	default:
		return nil
	}
}

//工厂B
type FactoryB struct {
}

func (this *FactoryB) Producepen(pentype string) pen {
	switch pentype {
	case "brush":
		return new(brushPen)
	default:
		return nil
	}
}
func (*FactoryB) Produceball(balltype string) ball {
	switch balltype {
	case "football":
		return new(footBall)
	default:
		return nil
	}
}
