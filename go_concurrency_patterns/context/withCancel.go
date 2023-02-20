package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	time.AfterFunc(time.Second*5, cancelFunc)
	go printNumber(ctx)
}

func printNumber(ctx context.Context) {

}
