# emptyCtx
一个emptyCtx永远不会被取消的,没有values,并且没有deadline,它不是一个结构体,因为一个emptyCtx在每次使用时必须有不同的地址。
```go
// emptyCtx 的原始类型是int类型
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
return
}

func (*emptyCtx) Done() <-chan struct{} {
return nil
}

func (*emptyCtx) Err() error {
return nil
}

func (*emptyCtx) Value(key any) any {
return nil
}

func (e *emptyCtx) String() string {
    switch e {
    case background:
        return "context.Background"
    case todo:
        return "context.TODO"
    }
    return "unknown empty Context"
}

var (
    background = new(emptyCtx)
    todo       = new(emptyCtx)
)

// Background()返回非nil的emptyCtx。
// 它从未被取消，没有价值，也没有截止日期。
// 它通常用于主函数、初始化和测试，并作为传入请求的顶级上下文。
func Background() Context {
	return background
}

// TODO返回非nil的空上下文。
// 当不清楚要使用哪个context或它还不可用时（因为周围函数尚未扩展为接受context参数），代码应该使用context.TTODO。
func TODO() Context {
	return todo
}
```