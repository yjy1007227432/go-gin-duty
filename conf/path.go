package main

import (
	"fmt"
	"go-gin-duty-master/pkg/gredis"
)

func main() {
	gredis.Setup()
	gredis.Set("yjy", "111", 60)
	str, _ := gredis.Get("yjy")
	fmt.Print(string(str))
}
