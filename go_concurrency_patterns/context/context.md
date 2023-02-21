# context.Context
context.Context interface 源码
```go
// Context 携带截止日期、取消信号和跨API边界的其他值。
// Context 的方法是并发安全的，可以由多个goroutine同时调用
type Context interface {
	Deadline() (deadline time.Time,ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
```
- Deadline()(deadline time.Time,ok bool):
  - Deadline返回代表此上下文完成的工作应被取消的时间。如果未设置截止日期，则Deadline返回ok==false。连续调用Deadline返回相同的结果。
- Done() <-chan struct{}:
  - Done返回一个通道，该通道在代表此上下文完成的工作应取消时关闭。如果此上下文永远无法取消，则Done可能返回nil。对Done的连续调用返回相同的值。在cancel函数返回后，Done通道可能会异步关闭。
  - WithCancel安排在调用cancel时关闭Done；WithDeadline安排在截止日期到期时关闭Done；WithTimeout安排在超时结束时关闭Done。
  - Done用于select语句：
```go
	//  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
```
- Err() error
  - Err返回非nil错误后，对Err的连续调用将返回相同的错误。
  - 如果Done尚未关闭，Err将返回nil。
  - 如果Done已关闭，Err将返回一个非nil错误，解释原因：
    - Canceled,如果上下文已取消
    - DeadlineExceeded,如果上下文的最后期限已过。
```go
var Canceled = errors.New("context canceled")

var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string   { return "context deadline exceeded" }
func (deadlineExceededError) Timeout() bool   { return true }
func (deadlineExceededError) Temporary() bool { return true }
```
- Value(key any) any
  - Value返回与key的上下文关联的值，如果没有值与key关联，则返回nil。使用相同的key连续调用Value返回相同的结果。