package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel() // 防止任务比超时时间短导致资源未释放
	time.Sleep(time.Second * 2)
	fmt.Println(ctx.Done())
	fmt.Println(ctx.Err())
}
