package main

import "fmt"

/*
	哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。
	一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别对应key和value。map中所有的key都有相同的类型，所以的value也有着相同的类型，但是key和value之间可以是
不同的数据类型。其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。
	所有这些操作（下标访问，delete）是安全的，即使这些元素不在map中也没有关系；如果一个查找失败将返回value类型对应的零值。
	map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作。禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址
无效。
	Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。这是故意的，每次都使用随机的遍历顺序
可以强制要求程序不会依赖具体的哈希函数实现。如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。
	map类型的零值是nil，也就是没有引用任何哈希表。
	map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常。
	通过key作为索引下标来访问map将产生一个value。如果key在map中是存在的，那么将得到与key对应的value；如果key不存在，那么将得到value对应类型的零值。
	map的下标语法将产生两个值；第二个是一个布尔值，用于报告元素是否真的存在。
	和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。要判断两个map是否包含相同的key和value，我们必须通过一个循环实现。
	有时候我们需要一个map或set的key是slice类型，但是map的key必须是可比较的类型，但是slice并不满足这个条件。不过，我们可以通过两个步骤绕过这个限制。第一步，定义
一个辅助函数k，将slice转为map对应的string类型的key，确保只有x和y相等时k(x) == k(y)才成立。然后创建一个key为string类型的map，在每次对map操作时先用k辅助函数将
slice转化为string类型。
	Map的value类型也可以是一个聚合类型，比如是一个map或slice。
*/
func main() {
	//	内置的make函数可以创建一个map
	ages := make(map[string]int)
	//也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value
	ages = map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	//Map中的元素通过key对应的下标语法访问
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // "32"
	//使用内置的delete函数可以删除元素
	delete(ages, "alice") // remove element ages["alice"]
	delete(ages, "alice")
	fmt.Println(ages)	//map[charlie:34]
	//x += y和x++等简短赋值语法也可以用在map上
	ages["bob"] += 1
	//不能对map的元素进行取址操作
	//_ = &ages["bob"] // compile error: cannot take address of map element

	//区分一个已经存在的0，和不存在而返回零值的0
	age, ok := ages["bob"]
	if !ok { /* "bob" is not a key in this map; age == 0. */ }
}
