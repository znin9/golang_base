package main

import (
	"fmt"
	"runtime"
	"time"
)

// 终结器实现同步
func main() {
	a := struct {
		name string
	}{name: "zs"}

	runtime.SetFinalizer(a, func(a struct{ name string }) { fmt.Println("被收集") })
	a = struct{ name string }{name: "lisi"}
	// runtime.GC()

	time.Sleep(time.Second * 5)
}
