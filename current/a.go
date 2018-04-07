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
	D int
	P string
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
		c.D = readInt()
		c.P = readString()
		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		fmt.Printf("Case #%d\n", i+1)
		res, ok := solve(c.D, c.P)
		if !ok {
			fmt.Fprintf(output, "Case #%d: IMPOSSIBLE\n", i+1)
		} else {
			fmt.Fprintf(output, "Case #%d: %d\n", i+1, res)
		}
	}
}

func solve(D int, P string) (int, bool) {
	fmt.Printf("solving %d %q\n", D, P)
	fmt.Println("power:", power(P))

	if D < numberOfS(P) {
		return -1, false
	}

	return 0, true
}

func power(P string) int {
	power := 1
	damages := 0
	for _, p := range P {
		switch p {
		case 'S':
			damages += power
		case 'C':
			power *= 2
		}
	}

	return damages
}

func numberOfS(P string) int {
	n := 0

	for _, p := range P {
		switch p {
		case 'S':
			n++
		}
	}

	return n
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
