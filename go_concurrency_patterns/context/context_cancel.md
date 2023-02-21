# cancelCtx
```go
// CancelFunc告诉操作放弃其工作。
// CancelFunc不会等待工作停止。
// 多个goroutine可以同时调用CancelFunc。
// 在第一次调用之后，对CancelFunc的后续调用不会执行任何操作。
type CancelFunc func()

// CancelCauseFunc的行为类似于CancelFunc，但还设置了取消原因。
// 可以通过在取消的上下文或其任何派生上下文上调用cause来检索此原因。
// 如果上下文已被取消，则CancelCauseFunc不会设置原因。
// 例如，如果childContext派生自parentContext：
// 如果在使用cause2取消childContext之前使用cause1取消parentContext，则Cause（parentContext）==Cause（childContext）==cause1
// 如果在使用cause1取消parentContext之前使用cause2取消childContext，则Cause（parentContext）==cause1和Cause（childContext）==couse2
type CancelCauseFunc func(cause error)
```
```go
// canceler是一种可以直接取消的上下文类型。实现是*cancelCtx和*timerCtx。
type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}
```
```go
// cancelCtx可以取消。取消时，它还会取消实现canceler的任何子级。
type cancelCtx struct {
	Context

	mu       sync.Mutex            // 保护以下字段
	done     atomic.Value          // chan结构｛｝的，延迟创建，由第一个cancel调用关闭
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
	cause    error                 // set to non-nil by the first cancel call
}
// methods

// implement Context interface Done() <- chan struct{}
func (c *cancelCtx) Done() <-chan struct{} {
    d := c.done.Load()
    if d != nil {
		return d.(chan struct{})
    }
    c.mu.Lock()
    defer c.mu.Unlock()
    d = c.done.Load()
    if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
    }
    return d.(chan struct{})
}

// implement Context interface Err() error 
func (c *cancelCtx) Err() error {
    c.mu.Lock()
    err := c.err
    c.mu.Unlock()
    return err
}

func (c *cancelCtx) Value(key any) any {
    if key == &cancelCtxKey {
		return c
    }
    return value(c.Context, key)
}
//
func (c *cancelCtx) String() string {
    return contextName(c.Context) + ".WithCancel"
}
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
    if err == nil {
        panic("context: internal error: missing cancel error")
    }
    if cause == nil {
        cause = err
    }
    c.mu.Lock()
    if c.err != nil {
        c.mu.Unlock()
        return // already canceled
    }
    c.err = err
    c.cause = cause
    d, _ := c.done.Load().(chan struct{})
    if d == nil {
        c.done.Store(closedchan)
    } else {
        close(d)
    }
    for child := range c.children {
    // NOTE: acquiring the child's lock while holding parent's lock.
        child.cancel(false, err, cause)
    }
    c.children = nil
    c.mu.Unlock()
    
    if removeFromParent {
        removeChild(c.Context, c)
    }
}
```