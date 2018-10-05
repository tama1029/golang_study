// +build ignore

package main

import (
	"fmt"
	"math"

	"github.com/d4l3k/go-pry/pry"
)

// type

type Abser interface {
	Abs() float64
}

type MyFloat2 float64

type Vertex2 struct {
	X, Y float64
}

// methods

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex2{3, 4}

	a = f // a MyFloat2 implements Abser
	fmt.Println(a)

	pry.Pry()
	a = v // a *Vertex2 implements Abser

	// In the following line, v is a Vertex2 (not *Vertex2)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

// 型を変えることで　既存のモジュールの関数の中身を置き換えることすら可能
