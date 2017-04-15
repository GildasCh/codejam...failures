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
	Hd int
	Ad int
	Hk int
	Ak int
	B  int
	D  int
}

func main() {
	sample := os.Args[1]
	fileIn := sample + ".in"
	fileOut := sample + ".out"

	var err error
	input, err = os.Open(fileIn)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", fileIn, err))
	}
	output, err = os.Create(fileOut)
	if err != nil {
		panic(fmt.Sprintf("creating %s: %v", fileOut, err))
	}
	defer input.Close()
	defer output.Close()

	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.Hd = readInt()
		c.Ad = readInt()
		c.Hk = readInt()
		c.Ak = readInt()
		c.B = readInt()
		c.D = readInt()

		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		s := solve(c.Hd, c.Hd, c.Ad, c.Hk, c.Ak, c.B, c.D, 0)
		if s == -1 {
			fmt.Fprintf(output, "Case #%d: IMPOSSIBLE\n", i+1)
		} else {
			fmt.Fprintf(output, "Case #%d: %d\n", i+1, s)
		}
	}
}

func solve(
	HdInitial int,
	Hd int,
	Ad int,
	Hk int,
	Ak int,
	B int,
	D int, counter int) int {

	if counter > 15 {
		return -1
	}

	if Hk <= 0 {
		return counter
	}

	if counter > 0 {
		Hd -= Ak
	}

	if Hd <= 0 {
		return -1
	}

	// Attack
	a := solve(HdInitial, Hd, Ad, Hk-Ad, Ak, B, D, counter+1)
	// Buff
	b := solve(HdInitial, Hd, Ad+B, Hk, Ak, B, D, counter+1)
	// Cure
	c := solve(HdInitial, HdInitial, Ad, Hk, Ak, B, D, counter+1)
	// Debuff
	d := solve(HdInitial, Hd, Ad, Hk, Ak-D, B, D, counter+1)

	return min(a, b, c, d)
}

func min(a, b, c, d int) int {
	ret := a
	if b != -1 && (ret == -1 || b < ret) {
		ret = b
	}
	if c != -1 && (ret == -1 || c < ret) {
		ret = c
	}
	if d != -1 && (ret == -1 || d < ret) {
		ret = d
	}
	return ret
}

func readInt() int {
	var i int
	fmt.Fscanf(input, "%d", &i)
	return i
}

func readString() string {
	var str string
	fmt.Fscanf(input, "%s", &str)
	return str
}

func readFloat() float64 {
	var x float64
	fmt.Fscanf(input, "%f", &x)
	return x
}

func readRune() rune {
	var x rune
	fmt.Fscanf(input, "%c", &x)
	return x
}
