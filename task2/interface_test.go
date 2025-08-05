

package main
import (
	"fmt"
	"math"
)

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
// 实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64{
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64{
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.radius
}

func main(){
	rect := Rectangle{
		length: 20, width:20
	}
	circle := Circle{
		radius: 5
	}

	shapes := []Shape{
		rect, circle
	}

	for _, shape := range shapes {
		fmt.Printf("%T, Area:%.2f, Perimeter:%.2f \n", shape, shape.Area(), shape.Perimeter())
	}

	fmt.Printf("\nRectangle: Area=%.2f, Perimeter=%.2f\n", rect.Area(), rect.Perimeter())
	fmt.Printf("Circle: Area=%.2f, Perimeter=%.2f\n", circle.Area(), circle.Perimeter())
}