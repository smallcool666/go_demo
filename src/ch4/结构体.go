package main

import (
	"fmt"
	"time"
)

/*
	结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。
	结构体赋值给变量之后，所有的成员也同样是变量，可以直接对每个成员赋值，或者是对成员取地址，然后通过指针访问。
	结构体成员的输入顺序也有重要的意义。成员一样，但是顺序不一样的结构体是不同的结构体类型。
	如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员。
	一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。（该限制同样适应于数组。）但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，
比如链表和树结构等。
	结构体类型的零值是每个成员都对是零值。通常会将零值作为最合理的默认值。
	如果结构体没有任何成员的话就是空结构体，写作struct{}。它的大小为0，也不包含任何信息，但是有时候依然是有价值的。

	结构体值也可以用结构体面值表示，结构体面值可以指定每个成员的值。
	以成员名字和相应的值来初始化，可以包含部分或全部的成员。在这种形式的结构体面值写法中，如果成员被忽略的话将默认用零值。因为，提供了成员的名字，所有成员出现的顺序并不重要。
	结构体可以作为函数的参数和返回值。考虑效率的话，较大的结构体通常会用指针的方式传入和返回。
	如果要在函数内部修改结构体成员的话，用指针传入是必须的；因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。

	如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较。相等比较运算符==将比较两个结构体的每个成员。
	可比较的结构体类型和其他可比较的类型一样，可以用于map的key类型。

	Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。
	得益于匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径。
	匿名成员都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员部分。
	结构体字面值并没有简短表示匿名成员的语法，结构体字面值必须遵循形状类型声明时的结构。
	Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。
	因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。同时，因为成员的名字是由其类型隐式地决定的，所有匿名成员也有可见性的规则约束。
	简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法。实际上，外层的结构体不仅仅是获得了匿名成员类型的所有成员，而且也获得了该类型导出的全部的方法。这个机制可以
用于将一个有简单行为的对象组合成有复杂行为的对象。组合是Go语言中面向对象编程的核心
*/

//公司的员工信息的结构体
type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}
var dilbert Employee

//通常一行对应一个结构体成员，成员的名字在前类型在后，不过如果相邻的成员类型如果相同的话可以被合并到一行
type Employee1 struct {
	ID int
	Name, Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

type Point struct {
	X, Y int
}
//常规的结构体，但是访问成员很繁琐
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}


func main() {
	//dilbert是一个变量，它所有的成员也同样是变量，我们可以直接对每个成员赋值
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	//点操作符也可以和指向结构体的指针一起工作
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"

	//结构体面值可以指定每个成员的值
	type Point struct{ X, Y int }
	var _ = Point{1, 2}
	//第二种写法，以成员名字和相应的值来初始化，可以包含部分或全部的成员
	var _ = Point{Y: 2}

	//结构体比较
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q) // "false"

	//使用匿名成员的结构体
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}
	//直接访问叶子属性而不需要给出完整的路径
	var w Wheel
	w.X = 8 // equivalent to w.Circle.Point.X = 8  可以通过匿名成员类型的名字来访问匿名成员的属性，这个是可选的
	w.Y = 8 // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20

	//结构体字面值并没有简短表示匿名成员的语法
	//w = Wheel{8, 8, 5, 20} // compile error: unknown fields
	//w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	//结构体字面值必须遵循形状类型声明时的结构
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	//Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。
	fmt.Printf("%#v\n", w)
	//Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}
}
