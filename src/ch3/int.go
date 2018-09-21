package main

import "fmt"

/*
	Go语言同时提供了有符号和无符号类型的整数运算。这里有int8、int16、int32和int64四种截然不同大小的有符号整形数类型，分别对应8、16、32、64bit大小的有符号整形数，与此对应的是uint8、
uint16、uint32和uint64四种无符号整形数类型。
	还有两种一般对应特定CPU平台机器字大小的有符号和无符号整数int和uint；其中int是应用最广泛的数值类型。这两种类型都有同样的大小，32或64bit，但是我们不能对此做任何的假设；因为不同的编
译器即使在相同的硬件平台上可能产生不同的大小。
	Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。这两个名称可以互换使用。
	同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。
	最后，还有一种无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程是才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。

	不管它们的具体大小，int、uint和uintptr是不同类型的兄弟类型。其中int和int32也是不同的类型，即使int的大小也是32bit，在需要将int当作int32类型的地方需要一个显式的类型转换操作，反之亦然。
	有符号整数采用2的补码形式表示，也就是最高bit位用作表示符号位，一个n-bit的有符号数的值域是从-2^{n-1}到 2^{n-1}−1。
	Go语言中关于算术运算、逻辑运算和比较运算的二元运算符，它们按照先级递减的顺序的排列：
		* / % << >> & &^
		+ - | ^
		== != < <= > >=
		&&
		||
	算术运算符+、-、*和/可以适用与于整数、浮点数和复数，但是取模运算符%仅用于整数间的运算。
	对于不同编程语言，%取模运算的行为可能并不相同。在Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此-5%3和-5%-3结果都是-2。
	除法运算符/的行为则依赖于操作数是否为全为整数，比如5.0/4.0的结果是1.25，但是5/4的结果是1，因为整数除法会向着0方向截断余数。

	如果原始的数值是有符号类型，而且最左边的bit为是1的话，那么最终结果可能是负的

	***整数、浮点数和字符串可以根据比较结果排序。许多其它类型的值可能是不可比较的，因此也就可能是不可排序的。

	bit位操作运算符，前面4个操作运算符并不区分是有符号还是无符号数:
		& 位运算 AND
		| 位运算 OR
		^ 位运算 XOR
		&^ 位清空 (AND NOT)
		<< 左移
		>> 右移
	位操作运算符^作为二元运算符时是按位异或（XOR），当用作一元运算符时表示按位取反
	位操作运算符&^用于按位置零（AND NOT）
	左移运算用零填充右边空缺的bit位，无符号数的右移运算也是用0填充左边空缺的bit位，但是有符号数的右移运算会用符号位的值填充左边空缺的bit位。因为这个原因，最好用无符号运算，这样你可以将整
数完全当作一个bit位模式处理。

	***无符号数往往只有在位运算或其它特殊的运算场景才会使用，就像bit集合、分析二进制文件格式或者是哈希和加密操作等。它们通常并不用于仅仅是表达非负数量的场合。

	许多整形数之间的相互转换并不会改变数值；它们只是告诉编译器如何解释这个值。但是对于将一个大尺寸的整数类型转为一个小尺寸的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度
	浮点数到整数的转换将丢失任何小数部分，然后向数轴零方向截断。应该避免对可能会超出目标类型表示范围的数值类型转换，因为截断的行为可能依赖于具体的实现
*/
func main() {
	//%取模运算符的符号和被取模数的符号总是一致的
	fmt.Println(5%3, -5%3, 5%-3, -5%-3) //2 -2 2 -2
	//整数除法会向着0方向截断余数
	fmt.Println(5/4, 5.0/4, 5.0/4.0, 5/4.0)//1 1.25 1.25 1.25

	//如果原始的数值是有符号类型，而且最左边的bit为是1的话，那么最终结果可能是负的
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"
	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

	//位运算演示
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // "00100010"
	fmt.Printf("%08b\n", y)    // "00000110"
	fmt.Printf("%08b\n", x|y)  // "00100110"
	fmt.Printf("%08b\n", x^y)  // "00100100"
	fmt.Printf("%08b\n", x&^y) // "00100000" 按位置零；按照y的bit位把x对应的bit位置零；y哪个bit位是1，就把x对应的bit位置为0
	fmt.Printf("%08b\n", x<<1) // "01000100"
	fmt.Printf("%08b\n", x>>1) // "00010001"

	//类型转换丢失精度
	f := 3.141 // a float64
	in := int(f)
	fmt.Println(f, in) // "3.141 3"

	//使用fmt包打印一个数值时，我们可以用%d、%o或%x参数控制输出的进制格式
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	xx := int64(0xdeadbeef)
	//fmt的两个使用技巧。通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。
	// 第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", xx)// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}
