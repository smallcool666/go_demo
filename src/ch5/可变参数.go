package main

import "fmt"

/*
	参数数量可变的函数称为为可变参数函数。典型的例子就是fmt.Printf和类似函数。Printf首先接收一个的必备参数，之后接收任意个数的后续参数。
	在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号“...”，这表示该函数会接收任意数量的该类型参数。
	如果原始参数已经是切片类型，我们该如何传递给sum？只需在最后一个参数后加上省略符。
	虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参数函数和以切片作为参数的函数是不同的。
	函数名的后缀f是一种通用的命名规范，代表该可变参数函数可以接收Printf风格的格式化字符串。
*/

//sum函数返回任意个int型参数的和
func sum(vals...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func f1(...int) {}
func g([]int) {}
func main() {
	fmt.Println(sum()) // "0"
	fmt.Println(sum(3)) // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"

	//向可变参数函数传递切片参数，切片后加...
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"

	//虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参数函数和以切片作为参数的函数是不同的。
	fmt.Printf("%T\n", f1) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"
}
