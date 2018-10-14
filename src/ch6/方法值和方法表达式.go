package main

import "fmt"

/*
	p.Distance叫作“选择器”，选择器会返回一个方法"值"->一个将方法(Point.Distance)绑定到特定接收器变量的函数。这个函数可以不通过指定其接收器即可被调用；即调用时不需要
指定接收器(译注：因为已经在前文中指定过了)，只要传入函数的参数即可。
	和方法"值"相关的还有方法表达式。当调用一个方法时，与调用一个普通的函数相比，我们必须要用选择器(p.Distance)语法来指定方法的接收器。
	当T是一个类型时，方法表达式可能会写作T.f或者(*T).f，会返回一个函数"值"，这种函数会将其第一个参数用作接收器，所以可以用通常(译注：不写选择器)的方式来对其进行调用。
	当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。
*/
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance // method value
	fmt.Println(distanceFromP(q)) // "5"

	distance := Point.Distance // method expression
	fmt.Println(distance(p, q)) // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
	// 译注：这个Distance实际上是指定了Point对象为接收器的一个方法func (p Point) Distance()，
	// 但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数，
	// 即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
	// 看起来本书中函数和方法的区别是指有没有接收器，而不像其他语言那样是指有没有返回值。
}
