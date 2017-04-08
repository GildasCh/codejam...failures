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
	S []bool
	K int
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
	runeLoop:
		for {
			r := readRune()
			switch r {
			case '+':
				c.S = append(c.S, true)
			case '-':
				c.S = append(c.S, false)
			default:
				break runeLoop
			}
		}
		c.K = readInt()
		C = append(C, c)
	}

	for i, c := range C {
		s := solve(c.S, c.K)
		if s >= 0 {
			fmt.Fprintf(output, "Case #%d: %d\n", i+1, s)
		} else {
			fmt.Fprintf(output, "Case #%d: IMPOSSIBLE\n", i+1)
		}
	}
}

func solve(S []bool, K int) int {
	// fmt.Println(S, K)

	counter := 0
	for {
		i := firstUnhappy(S)
		if i == -1 {
			return counter
		}

		ok := flip(S, i, K)
		if !ok {
			return -1
		}

		// fmt.Println("Flip OK:", S)
		counter++
	}

	panic("Should not get here")
	return -2
}

func firstUnhappy(S []bool) int {
	for i, s := range S {
		if !s {
			return i
		}
	}
	return -1
}

func flip(S []bool, i int, K int) bool {
	if i+K > len(S) {
		return false
	}

	for k := 0; k < K; k++ {
		S[i+k] = !S[i+k]
	}

	return true
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
