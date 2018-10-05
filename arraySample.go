// +build ignore

package main

import (
	"fmt"
)

func arrayTest() {
	var arr [10]int
	arr[0] = 10
	arr[0] = 11
	fmt.Println(arr)
}

func minTest() {
	a := [10]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	doubleArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(doubleArray)
}

func main() {
	arrayTest()
	minTest()
}
