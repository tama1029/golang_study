// +build ignore

package main

import (
	"fmt"
)

func sliceTest() {
	var fslice []int
	fmt.Println(fslice)

	slice := []byte{'a', 'b', 'c'}
	fmt.Println(slice[:])
	fmt.Println(slice[1:2]) // 1だけ
	fmt.Println(slice[:2])  // 0, 1だけ
	slice[1] = 3
	fmt.Println(slice[:1])  // 0だけ
	fmt.Println(slice[1:2]) // 1だけ
	fmt.Println(slice[1:2]) // 1だけ
	fmt.Println(slice[1:2]) // 1だけ
	fmt.Println(slice[1:2]) // 1だけ
	fmt.Println(slice[:2])  // 0, 1だけ
}

func mapTest() {
	rating := map[string]float32{"C#": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// mapは２つの戻り値があります。２つ目の戻り値では、もしkeyが存在しなければ、okはfalseに、存在すればokはtrueになります。
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
}
func main() {
	sliceTest()
	mapTest()
}
