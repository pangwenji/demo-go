package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func PrintShapeInfo(s Shape) {
	fmt.Println("画面积", s.Area())
	fmt.Println("画周长", s.Perimeter())
}

func main() {
	// 创建 Circle 对象
	circle := &Circle{Radius: 5}
	fmt.Println("Circle:")
	PrintShapeInfo(circle) // 使用接口方法

	// 创建 Rectangle 对象
	rectangle := &Rectangle{Width: 4, Height: 6}
	fmt.Println("\nRectangle:")
	PrintShapeInfo(rectangle) // 使用接口方法

}
