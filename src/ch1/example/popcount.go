package main
/*
	代码定义了一个PopCount函数，用于返回一个数字中含二进制1bit的个数。它使用init初始化函数来生成辅助表格pc，pc表格用于处理每个8bit宽度的数字含二进制的1bit的bit个数，这样
的话在处理64bit宽度的数字时就没有必要循环64次，只需要8次查表就可以了。
	演示init函数的用法，并且演示了如何预生成辅助表格。
*/
import "fmt"

// pc[i] is the population count of i.
var pc [256]byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func main() {
	arr := []uint64{19999, 100, 256}
	for _, v := range arr {
		fmt.Println(v, "  ", PopCount(v))
	}
}