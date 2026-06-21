package main

import "fmt"

type Shape interface{
	Area() float64
	Perimetr() float64
}

type Rectangle struct{
	Width float64
	Height float64
}

func CreateRectangle(W float64, H float64) Rectangle{
	return Rectangle{
		Width: W,
		Height: H,
	}
}

func (r Rectangle) Area() float64{
	return r.Height * r.Width
}

func (r Rectangle) Perimetr() float64{
	return 2 *(r.Height + r.Width)
}

type Circle struct{
	Radius float64
}
func CreateCircle(R float64) Circle{
	return Circle{
		Radius: R,
	}
}
func (c Circle) Area() float64{
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimetr()float64{
	return 2 * 3.14 * c.Radius
}

func main(){
	govno := make([]Shape, 0)
	govno = append(govno, CreateCircle(4))
	govno = append(govno, CreateRectangle(3, 4))
	fmt.Println("Площадь, периметр круга: ", govno[0].Area(), govno[0].Perimetr())
	fmt.Println("Площадь, периметр прямоугольника", govno[1].Area(), govno[1].Perimetr())
}
