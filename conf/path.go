package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "姚俊毅"
	old := "姚俊毅"
	new := "呵呵呵"
	// n < 0 ,用 new 替换所有匹配上的 old；n=-1:  3915abcd3915abcd3915abcd3915abcd3915abcd
	fmt.Println("n=-1: ", strings.Replace(s, old, new, -1))

	// n = 0 ,不替换任何匹配的 old; n=0: 123abcd123abcd123abcd123abcd123abcd
	fmt.Println("n=0: ", strings.Replace(s, old, new, 0))

	// n = 1 ,用 new 替换第一个匹配的 old；n=-1:  3915abcd123abcd123abcd123abcd123abcd
	fmt.Println("n=1: ", strings.Replace(s, old, new, 1))

	// n = 2 ,用 new 替换第二个匹配的 old；n=-1:  3915abcd3915abcd123abcd123abcd123abcd
	fmt.Println("n=0: ", strings.Replace(s, old, new, 2))

	// n = 2,old="" 在最前面插入二个new；n=2:  39151391523abcd123abcd123abcd123abcd123abcd
	fmt.Println("n=2: ", strings.Replace(s, "", new, 2))
}
