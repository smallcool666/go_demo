package main

import (
	"fmt"
	"sort"
)

/*
	sort包内置的提供了根据一些排序函数来对任何序列排序的功能。它的设计非常独到。在很多语言中，排序算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。相比之下，
Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实
现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。
	一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是sort.Interface的三个方法：Len(), Less(), Swap()。
	对字符串切片的排序是很常用的需要，所以sort包提供了StringSlice类型，也提供了Strings函数能让上面这些调用简化成sort.Strings(names)。
	sort函数会交换很多对元素，所以如果每个元素都是指针会更快。
	sort包中提供了Reverse函数将排序顺序转换成逆序，可以对升序的方法使用Reverse即可得到降序的结果。
	sort包定义了一个不公开的struct类型reverse，它嵌入了一个sort.Interface。reverse的Less方法调用了内嵌的sort.Interface值的Less方法，但是通过交换索引的方式使排序结果变成逆序。
	实现了sort.Interface的具体类型不一定是切片类型。
	尽管对长度为n的序列排序需要 O(n log n)次比较操作，检查一个序列是否已经有序至少需要n−1次比较。sort包中的IsSorted函数帮我们做这样的检查。像sort.Sort一样，它也使用sort.Interface
对这个序列和它的排序函数进行抽象，但是它从不会调用Swap方法。
	为了使用方便，sort包为[]int,[]string和[]float64的正常排序提供了特定版本的函数和类型。对于其他类型，例如[]int64或者[]uint，尽管路径也很简单，还是依赖我们自己实现。
*/

type StringSlice []string
func (p StringSlice) Len() int { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func main() {
	names := []string{"aaa", "abc", "dgda"}
	sort.Sort(StringSlice(names))

	//判断切片是否有序
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values) // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values) // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
}
