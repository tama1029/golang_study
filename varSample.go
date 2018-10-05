// +build ignore

package main

import (
	"fmt"
	"strconv"
)

var x = 10

const (
	TestInteger = 3
	testString  = "hogee"
)

const (
	enum1 = iota
	enum2
	enum3
)

func square(x int) int {
	vhoge := 2
	return vhoge*x*x + TestInteger
}

func aaa(x int) string {
	var available0 bool
	var integertest int
	var float32test float32
	var stringtest string
	fmt.Println(available0)
	fmt.Println(integertest)
	fmt.Println(float32test)
	fmt.Println(stringtest)
	return "hogee" + strconv.Itoa(TestInteger)
}

func stringTest() {
	s := "stringtest"
	fmt.Println(s + testString)
}

func enumTest() {
	fmt.Println(enum1)
	fmt.Println(enum2)
	fmt.Println(enum3)
}

func main() {
	fmt.Println(x)
	fmt.Println(square(5))
	fmt.Println(aaa(3))
	stringTest()
	enumTest()
}
