package main

import (
	"fmt"
	"os"
)

var input *os.File
var output *os.File

var T int
var C []Case

type Case struct {
	R, B int
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.R, c.B = readInt(), readInt()
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		fmt.Fprintf(os.Stderr, "Case #%v\n", c)
		res := solve(c.R, c.B)
		fmt.Printf("Case #%d: %d\n", i+1, res)
	}
}

func solve(R, B int) int {
	if R < B {
		return solve(B, R)
	}

	res := 0

	b, u := best(R, 9999999)
	res += b
	R = u

	bs := 1

	for B >= bs && b > 0 {
		fmt.Fprintf(os.Stderr, "trying with bs=%d\n", bs)

		res++
		B -= bs

		if R > 0 {
			b, u = best(R, B/bs)
			res += b
			B -= b * bs
			R = u
		}

		bs++

		fmt.Fprintf(os.Stderr, "ended with bs=%d, B=%d, res=%d\n",
			bs, B, res)
	}

	return res
}

func max(r, b int) int {
	if r > b {
		return r
	}
	return b
}

func best(r, maxRes int) (res int, unused int) {
	unused = r
	for k := 1; k <= unused && res < maxRes; k++ {
		res = k
		unused -= k
	}

	fmt.Fprintf(os.Stderr, "best(%d,%d) return %d,%d\n",
		r, maxRes, res, unused)
	return
}

func readInt() int {
	var i int
	fmt.Fscanf(os.Stdin, "%d", &i)
	return i
}

func readString() string {
	var str string
	fmt.Fscanf(os.Stdin, "%s", &str)
	return str
}

func readChar() rune {
	var c rune
	fmt.Fscanf(os.Stdin, "%c", &c)
	return c
}
