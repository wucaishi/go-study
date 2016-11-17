// DeferS
package main

import (
	"fmt"
)

func main() {
	//	for i := 0; i <= 3; i++ {
	//		defer fmt.Println(i)
	//	}
	fmt.Println(f())
}
func f() (result int) {
	defer func() {
		result++
	}()
	return 1
}
