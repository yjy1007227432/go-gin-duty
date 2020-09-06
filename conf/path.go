package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print(time.Now().Format("2006-01-02"))
	fmt.Println("2020-08-07" > "2020-09-06")
}
