package main

import (
	"fmt"
	"time"
)

func main() {
	nowDay := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	fmt.Print(nowDay)
}
