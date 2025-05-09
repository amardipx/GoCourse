package intermediate

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

type Shape struct {
	Rectangle
}

// Method with value reciever
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer reciever
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {

	rect := Rectangle{
		length: 10,
		width:  9,
	}

	area := rect.Area()
	fmt.Println("Area of the rectangle before scaling is:", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of the rectangle after scaling is:", area)

	num1 := MyInt(-5)
	num2 := MyInt(10)

	fmt.Println(num1.isPositive())
	fmt.Println(num2.isPositive())

	s := Shape{
		Rectangle: Rectangle{
			length: 10,
			width:  9,
		},
	}

	fmt.Println(s.Area())

}

// User defined type
type MyInt int

// function on a user defined type
func (m MyInt) isPositive() bool {
	return m > 0
}
