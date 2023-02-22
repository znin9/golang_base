package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"
)

func main() {
	// test http.Serve(l net.Listener,h Handler)
	// Serve()

	// test http.ListenerAndServe
	// 效果同上Serve()
	// ListenAndServe()

	// test 自己创建一个多路复用器
	// ListenAndServeWithServeMux()

	// test Server的Handler是一个HandlerFunc
	HandlerFuncInServer()
}

// Demo: http.Serve(l net.Listener,h Handler)

type MyHandler struct {
	count atomic.Int64
}

func (mh *MyHandler) increment() int {
	newVal := mh.count.Add(1)
	return int(newVal)
}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count := mh.increment()
	w.Write([]byte(strconv.Itoa(count)))
}

// Serve 此时监听的TCP端口发送过来的请求，所有请求都由MyHandler的 ServeHTTP这个处理程序所处理。
// 因为在本示例中Handler(即是MyHandler)并没有对请求URL进行pattern匹配再使用对应的HandlerFunc进行处理。
func Serve() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	http.Serve(l, new(MyHandler))
}

// ===========================================

// Demo: http.ListenerAndServe(addr string,h Handler)
// 此时任何路请求路径都由MyHandler提供的ServeHTTP方法进行处理

func ListenAndServe() {
	err := http.ListenAndServe(":8080", new(MyHandler))
	if err != nil {
		panic(err)
	}
}

// =========================================
// Demo: 使用自己创建的HTTP请求多路复用器,并且在多路复用器(ServeMux)上注册一个处理函数

func ListenAndServeWithServeMux() {
	// ServerMux 是一个HTTP请求多路复用器
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 打印该请求是用的哪个Handler
		_, pattern := mux.Handler(r)
		log.Printf("pattern=%s\n", pattern)
		w.Write([]byte(fmt.Sprintln(pattern)))
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

// =============================================
// Demo: 使用HandlerFunc装配在Server的Handler处，所有请求都由HandlerFunc进行处理

func HandlerFuncInServer() {

	f := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all http request handle by this HandlerFunc"))
	}

	err := http.ListenAndServe(":8080", f)
	if err != nil {
		panic(err)
	}
	// TODO
}

// ==============================================
