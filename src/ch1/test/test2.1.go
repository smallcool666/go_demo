package main
/*
	向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。
*/

import "fmt"
type Celsius float64
type Fahrenheit float64

type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	AbsoluteZeroK Kelvin = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

/*
	开尔文（英语：Kelvin）
	是温度的计量单位。它是国际单位制（SI）的七个基本单位之一，符号为K。以开尔文计量的温度标准称为热力学温标，其零点为绝对零度。在热力学的经典表述中，绝对零度下所有热运动停止。1开尔文定义为
水的三相点与绝对零度相差的1⁄273.16。[1]水的三相点是0.01°C，因此温度变化1摄氏度，相当于变化了1开尔文。

	与其它温度单位之间的转换主条目：温度单位换算
		从开氏温标换算至其他温度单位
		从其他温度单位换算至开氏温标
	摄氏温标
		[°C] = [K] − 273.15
		[K] = [°C] + 273.15
	华氏温标
		[°F] = [K] × 9⁄5 − 459.67
		[K] = ([°F] + 459.67) × 5⁄9
*/
func KToC(k Kelvin) Celsius {
	return Celsius(k -  273.15 )
}
func CToK(c Celsius) Kelvin {
	return Kelvin(c +  273.15 )
}
func main() {
	
}
