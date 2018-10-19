package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	对那些基于http.Handler接口的服务器API做更进一步的学习。
	http.Handler接口只有一个方法ServeHTTP。http包里面还有一个ListenAndServe函数，需要一个例如“localhost:8000”的服务器地址，和一个所有请求都可以分派的Handler接口实例。它会
一直运行，直到这个服务因为一个错误而失败（或者启动失败），它的返回值一	定是一个非空的错误。
	net/http包提供了一个请求多路器ServeMux来简化URL和handlers的联系。一个ServeMux将一批http.Handler聚集到一个单一的http.Handler中。
	语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，因为http.HandlerFunc是一个类型。
	HandlerFunc显示了在Go语言接口机制中一些不同寻常的特点。这是一个有实现了接口http.Handler方法的函数类型。ServeHTTP方法的行为调用了它本身的函数。因此HandlerFunc是一个让函数
值满足一个接口的适配器，这里函数和这个接口仅有的方法有相同的函数签名。
	因为handler通过这种方式注册非常普遍，ServeMux有一个方便的HandleFunc方法，它帮我们简化handler注册代码。
	net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。现在，为了使用DefaultServeMux作为服务器的主handler，我们不需要将它传给
ListenAndServe函数；nil值就可以工作。
	web服务器在一个新的协程中调用每一个handler，所以当handler获取其它协程或者这个handler本身的其它请求也可以访问的变量时一定要使用预防措施比如锁机制。
*/


func main() {
	db := database{"shoes": 50, "socks": 5}
	//监听指定端口，执行Handle接口实例db的ServerHTTP方法
	log.Fatal(http.ListenAndServe("localhost:8000", db))

	//创建一个ServeMux并且使用它将URL和相应处理/list和/price操作的handler联系起来
	mux := http.NewServeMux()
	//语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，因为http.HandlerFunc是一个类型。
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	//简化handler注册代码
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	//使用DefaultServeMux作为服务器的主handler
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
type dollars float32
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
type database map[string]dollars

//ServeHTTP方法的简单实现，只是但因map中的值
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
//简单的url路由
func (db database) ServeHTTP1(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request){}
func (db database) price(w http.ResponseWriter, req *http.Request){}


