/*
	使用slice实现栈的功能
*/
package main

import "fmt"

//删除slice中间的某个元素并保存原有的元素顺序，可以通过内置的copy函数将后面的子slice向前依次移动一位完成
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	stack := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(stack, 2)) // "[5 6 8 9]"
	stack = append(stack, 10) // push v
	top := stack[len(stack)-1] // top of stack
	stack = stack[:len(stack)-1] // pop
}
