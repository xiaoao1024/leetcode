package solution

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_reverse(t *testing.T) {

	var (
		a int
		b int32
		c int64
		d int8
		e string
		f bool
		g rune

		classA ClassA
		classB ClassB
	)
	fmt.Println("size of int:", unsafe.Sizeof(a))
	fmt.Println("size of int32:", unsafe.Sizeof(b))
	fmt.Println("size of int64:", unsafe.Sizeof(c))
	fmt.Println("size of int8:", unsafe.Sizeof(d))
	fmt.Println("size of string:", unsafe.Sizeof(e))
	fmt.Println("size of bool:", unsafe.Sizeof(f))
	fmt.Println("size of rune:", unsafe.Sizeof(g))

	fmt.Println("size of class a:", unsafe.Sizeof(classA))
	fmt.Println("size of class b:", unsafe.Sizeof(classB))
}

type ClassA struct {
	A int8
	B int64
	C int16
}

type ClassB struct {
	A int8
	B int16
	C int64
}
