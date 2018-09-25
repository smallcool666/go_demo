package main

import (
	"fmt"
	"strconv"
)

/*
	一个字符串是一个不可改变的字节序列。
	文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。
	内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目）。
	索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。
	第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节。
	子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（并不包含j本身）生成一个新字符串。生成的新字符串将包含j-i个字节。
	不管i还是j都可能被忽略，当它们被忽略时将采用0作为开始位置，采用len(s)作为结束的位置。
	+操作符将两个字符串链接构造一个新字符串。
	字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变，当然我们也可以给一个字符串变量分配一个新字符串值。
	因为字符串是不可修改的，因此尝试修改字符串内部数据的操作是被禁止的。
	***不变性意味如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的。同样，一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地
共享相同的内存，因此字符串切片操作代价也是低廉的。在这两种情况下都没有必要分配新的内存。


	字符串值也可以用字符串面值方式编写，只要将一系列字节序列包含在双引号即可。
	因为Go语言源文件总是用UTF8编码，并且Go语言的文本字符串也以UTF8编码的方式处理，因此我们可以将Unicode码点也写到字符串面值中。
	在一个双引号包含的字符串面值中，可以用以反斜杠\开头的转义序列插入任意的数据。
	一个原生的字符串面值形式是`...`，使用反引号代替双引号。在原生的字符串面值中，没有转义操作；全部的内容都是字面的意思，包含退格和换行，因此一个程序中的原生字符串面值
可能跨越多行（译注：在原生字符串面值内部是无法直接写字符的，可以用八进制或十六进制转义或+"```"链接字符串常量完成）。唯一的特殊处理是会删除回车以保证在所有平台上的值都
是一样的，包括那些把回车也放入文本文件的系统（译注：Windows系统会把回车和换行一起放入文本文件中）。


	Unicode（ http://unicode.org ）收集了这个世界上所有的符号系统，包括重音符号和其它变音符号，制表符和回车符，还有很多神秘的符号，每个符号都分配一个唯一的Unicode码点，
Unicode码点对应Go语言中的rune整数类型（译注：rune是int32等价类型）。
	UTF8是一个将Unicode码点编码为字节序列的变长编码。UTF8编码由Go语言之父Ken Thompson和Rob Pike共同发明的，现在已经是Unicode的标准。UTF8编码使用1到4个字节来表示每个
Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用2或3个字节表示。每个符号编码后第一个字节的高端bit位用于表示总共有多少编码个字节。如果第一个字节的高端bit为0，则
表示对应7bit的ASCII字符，ASCII字符每个字符依然是一个字节，和传统的ASCII编码兼容。如果第一个字节的高端bit是110，则说明需要2个字节；后续的每个高端bit都以10开头。更大的
Unicode码点也是采用类似的策略处理。
		0xxxxxxx runes 0-127 (ASCII)
		110xxxxx 10xxxxxx 128-2047 (values <128 unused)
		1110xxxx 10xxxxxx 10xxxxxx 2048-65535 (values <2048 unused)
		11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 65536-0x10ffff (other values unused)
	变长的编码无法直接通过索引来访问第n个字符，但是UTF8编码获得了很多额外的优点。首先UTF8编码比较紧凑，完全兼容ASCII码，并且可以自动同步：它可以通过向前回朔最多2个字节
就能确定当前字符编码的开始字节的位置。它也是一个前缀编码，所以当从左向右解码时不会有任何歧义也并不需要向前查看（译注：像GBK之类的编码，如果不知道起点位置则可能会出现歧
义）。没有任何字符的编码是其它字符编码的子串，或是其它编码序列的字串，因此搜索一个字符时只要搜索它的字节编码序列即可，不用担心前后的上下文会对搜索结果产生干扰。同时UTF8
编码的顺序和Unicode码点的顺序一致，因此可以直接排序UTF8编码序列。同时因为没有嵌入的NUL(0)字节，可以很好地兼容那些使用NUL作为字符串结尾的编程语言。
	Go语言字符串面值中的Unicode转义字符让我们可以通过Unicode码点输入特殊的字符。有两种形式：\uhhhh对应16bit的码点值，\Uhhhhhhhh对应32bit的码点值，其中h是一个十六进制数
字；一般很少需要使用32bit的形式。每一个对应码点的UTF8编码。
	对于小于256码点值可以写在一个十六进制转义字节中，例如'\x41'对应字符'A'，但是对于更大的码点则必须使用\u或\U转义形式。因此，'\xe4\xb8\x96'并不是一个合法的rune字符，虽
然这三个字节对应一个有效的UTF8编码的码点。
	得益于UTF8编码优良的设计，诸多字符串操作都不需要解码操作。
	Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。
	每一个UTF8字符解码，不管是显式地调用utf8.DecodeRuneInString解码或是在range循环中隐式地解码，如果遇到一个错误的UTF8编码输入，将生成一个特别的Unicode字符'\uFFFD'，在
印刷中这个符号通常是一个黑色六角或钻石形状，里面包含一个白色的问号（?）。当程序遇到这样的一个字符，通常是一个危险信号，说明输入并不是一个完美没有错误的UTF8字符串。
	UTF8字符串作为交换格式是非常方便的，但是在程序内部采用rune序列可能更方便，因为rune大小一致，支持数组索引和方便切割。
	string接受到[]rune的类型转换，可以将一个UTF8编码的字符串解码为Unicode字符序列
	如果是将一个[]rune类型的Unicode字符slice或数组转为string，则对它们进行UTF8编码
	将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串


	标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
	strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
	bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer
类型将会更有效。
	strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
	unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。
	一个字符串是包含的只读字节数组，一旦创建，是不可变的。相比之下，一个字节slice的元素则可以自由地修改。
	从概念上讲，一个[]byte(s)转换是分配了一个新的字节数组用于保存字符串数据的拷贝，然后引用这个底层的字节数组。编译器的优化可以避免在一些场景下分配和复制字符串数据，但总
的来说需要确保在变量b被修改的情况下，原始的s字符串也不会改变。将一个字节slice转到字符串的string(b)操作则是构造一个字符串拷贝，以确保s2字符串是只读的。
	为了避免转换中不必要的内存分配，bytes包和strings同时提供了许多实用函数。
	bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要处理化，因为
零值也是有效的。


	将一个整数转为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一个方法是用strconv.Itoa(“整数到ASCII”)
	FormatInt和FormatUint函数可以用不同的进制来格式化数字
	fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多
	如果要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无符号整数的ParseUint函数
*/
func main() {
	fmt.Println(len("hello,world!"))   //12
	fmt.Println(len("你好，世界"))      //15

	str := "hello, world"
	//子字符串操作s[i:j]
	fmt.Println(str[0:5]) // "hello"
	//子字符串操作s[i:j],i还是j使用默认值
	fmt.Println(str[:5]) // "hello"
	fmt.Println(str[7:]) // "world"
	fmt.Println(str[:]) // "hello, world"
	//+操作符连接两个字符串
	fmt.Println("goodbye" + str[5:]) // "goodbye, world"
	//尝试修改字符串内部数据的操作是被禁止的
	//str[0] = 'L' // compile error: cannot assign to s[0]

	//原生的字符串面值
	s := `go语言	"c语言"
php语言 'java语言' "Python" \n\r\t`
	fmt.Println(s)
	//go语言	"c语言"
	//php语言 'java语言' "Python" \n\r\t

	//Unicode转义字符让我们可以通过Unicode码点输入特殊的字符
	fmt.Println("\u4e16\u754c", "\xe4\xb8\x96\xe7\x95\x8c", "\U00004e16\U0000754c")    //世界 世界 世界

	//range循环在处理字符串的时候，会自动隐式解码UTF8字符串
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	/*
		0	'H'	72
		1	'e'	101
		2	'l'	108
		3	'l'	108
		4	'o'	111
		5	','	44
		6	' '	32
		7	'世'	19990
		10	'界'	30028
	*/

	//string转换成[]rune
	s = "プログラム"
	//Printf中的% x参数用于在每个十六进制数字前插入一个空格
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	//将一个[]rune类型的Unicode字符slice或数组转为string，则对它们进行UTF8编码
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	fmt.Println(string(65)) // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"

	//整数转字符串
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	//FormatInt和FormatUint函数可以用不同的进制来格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	//字符串解析为整数
	x, err := strconv.Atoi("123") // x is an int
	z, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(z, err)
}
