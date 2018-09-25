package main

import "fmt"

/*
	在原内存空间将[]int类型的slice反转
*/

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1,2,3,4,5,6,7,8,9}
	reverse(s)
	fmt.Println(s) //[9 8 7 6 5 4 3 2 1]

	//一种将slice元素循环向左镟转n个元素的方法是三次调用reverse反转函数，第一次是反转开头的n个元素，然后是反转剩下的元素，最后是反转整个slice的元素。
	s1 := []int{1,2,3,4,5,6,7,8,9}
	reverse(s1[:2])
	reverse(s1[2:])
	reverse(s1)
	fmt.Println(s1)		//[3 4 5 6 7 8 9 1 2]
}
