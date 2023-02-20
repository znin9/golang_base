package main

/*
*

	Mutex 公平:

	Mutex一共有两种操作模式
	在普通模式下waiters按照FIFO队列的方法进行获取锁，但是一个被唤醒的waiter
	并不一定会获取锁，因为会和新到来的goroutine进行锁竞争，
	由于一般新到来的goroutine都处于CPU上执行，更容易获取到Mutex。
	这时该waiter就会继续入队，但是当一个waiter超过1ms都未获取到Mutex的所有权时
	就会从普通模式变为饥饿模式。

	在饥饿模式下，队列总是按照FIFO队列的方式顺序获取Mutex的所有权。
*/
func main() {
}
