package main
import (
	"fmt"
	"strings"
)

type Struct struct {
	X int
	Y string
}

var (
	v1 = Struct{1, "v1"}   // has type Struct
	v2 = Struct{X: 1}      // Y:"" is implicit
	v3 = Struct{ }         // X:0 and Y:""
	pv  = &Struct{1, "pv"} // has type *Struct
)

func main() {
	// Pointer
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Struct
	fmt.Println(Struct{12, "Hello"})
	v := Struct{1, "Woo"}
	v.X = 4
	fmt.Println(v.X)

	po := &v
	po.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, pv, v2, v3)

	// Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slice ※スライスとは可変長の配列のこと
	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	// Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)

	// Slice defaults
	k := []int{2, 3, 5, 7, 11, 13}

	k = k[1:4]
	fmt.Println(k)

	k = k[:2]
	fmt.Println(k)

	k = k[1:]
	fmt.Println(k)

	// Slice length and capacity ※キャパシティとはそのスライスが保持できる要素数の最大値
	sa := []int{2, 3, 5, 7, 11, 13}
	printSlice(sa)

	// Slice the slice to give it zero length.
	sa = sa[:0]
	printSlice(sa)

	// Extend its length.
	sa = sa[:4]
	printSlice(sa)

	// Drop its first two values.
	sa = sa[2:]
	printSlice(sa)

	// Nil slices
	var ss []int
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("nil!")
	}

	// Creating a slice with make
	e := make([]int, 5)
	printSlice2("e", e)

	f := make([]int, 0, 5)
	printSlice2("f", f)

	g := f[:2]
	printSlice2("g", g)

	h := g[2:5]
	printSlice2("h", h)

	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	var as []int
	printSlice(as)
	// append works on nil slices.
	as = append(as, 0)
	printSlice(as)
	// The slice grows as needed.
	as = append(as, 1)
	printSlice(as)
	// We can add more than one element at a time.
	as = append(as, 2, 3, 4)
	printSlice(as)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}