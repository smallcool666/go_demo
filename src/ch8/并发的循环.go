package main

import (
	"log"
	"os"
	"sync"
)

/*
	子问题都是完全彼此独立的问题被叫做易并行问题(译注：embarrassingly parallel，直译的话更像是尴尬并行)。
	没有什么直接的办法能够等待goroutine完成，但是我们可以改变goroutine里的代码让其能够将完成情况报告给外部的goroutine知晓，使用的方式是向一个共享的channel中发送事件。因为我们已经知道内
部的goroutine只有len(filenames)，所以外部的goroutine只需要在返回之前对这些事件计数。
	为了知道最后一个goroutine什么时候结束(最后一个结束并不一定是最后一个开始)，我们需要一个递增的计数器，在每一个goroutine启动时加一，在goroutine退出时减一。这需要一种特殊的计数器，这个
计数器需要在多个goroutine操作时做到安全并且提供提供在其减为零之前一直等待的一种方法。这种计数类型被称为sync.WaitGroup。
	Add是为计数器加一，必须在worker goroutine开始之前调用，而不是在goroutine中；否则的话我们没办法确定Add是在"closer" goroutine调用Wait之前被调用。并且Add还有一个参数，但Done却没有任
何参数；其实它和Add(-1)是等价的。我们使用defer来确保计数器即使是在出错的情况下依然能够正确地被减掉。
	观察一下我们是怎样创建一个closer goroutine，并让其等待worker们在关闭掉sizes channel之前退出的。两步操作：wait和close，必须是基于sizes的循环的并发。

	无穷无尽地并行化并不是什么好事情，因为不管怎么说，你的系统总是会有一个些限制因素，比如CPU核心数会限制你的计算负载，比如你的硬盘转轴和磁头数限制了你的本地磁盘IO操作频率，比如你的网络
带宽限制了你的下载速度上限，或者是你的一个web服务的服务容量上限等等。为了解决这个问题，我们可以限制并发程序所使用的资源来使之适应自己的运行环境。
	我们可以用一个有容量限制的buffered channel来控制并发，这类似于操作系统里的计数信号量概念。从概念上讲，channel里的n个空槽代表n个可以处理内容的token(通行证)，从channel里接收一个值会释
放其中的一个token，并且生成一个新的空槽位。这样保证了在没有接收介入时最多有n个发送操作。
*/

// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore errorsizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}