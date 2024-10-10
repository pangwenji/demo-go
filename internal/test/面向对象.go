package main

import "fmt"

type Car struct {
	Speed int
	Brand string
}

func (c *Car) Accelerate(increate int) {
	c.Speed += increate
	fmt.Printf("汽车达到每秒钟 km/h\n", c.Brand, c.Speed)
}

func (c *Car) broker(decreate int) {
	c.Speed -= decreate
	if c.Speed == 0 {
		c.Speed = 0
	}
	fmt.Printf("汽车减速km/h\n", c.Brand, c.Speed)
}

type ElectricCar struct {
	Car
	Battery int
}

func (e *ElectricCar) Change(amount int) {
	e.Battery += amount
	if e.Battery > 100 {
		e.Battery = 100
	}
	fmt.Printf("%s 的电池充电到 %d%%\n", e.Brand, e.Battery)
}

func main() {
	c := Car{
		Speed: 0,
		Brand: "小米",
	}

	c.Accelerate(50)
	c.broker(20)

	e := ElectricCar{
		Car: Car{
			Brand: "哈哈",
			Speed: 34,
		},
		Battery: 50,
	}

	e.Brand = "telse"
	e.Change(65)
	fmt.Printf("电车的速度是", e.Brand, e.Speed)
}
