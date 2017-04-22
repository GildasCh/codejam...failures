package main

import (
	"fmt"
	"os"
)

var input *os.File
var output *os.File

var T int64
var C []Case

type Case struct {
	D int64
	N int64
	K []int64
	S []int64
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

	for i := int64(0); i < T; i++ {
		c := Case{}
		c.D = readInt()
		c.N = readInt()
		c.K = make([]int64, c.N)
		c.S = make([]int64, c.N)
		for i := int64(0); i < c.N; i++ {
			c.K[i] = readInt()
			c.S[i] = readInt()
		}
		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		ret := solve(c.D, c.N, c.K, c.S)
		fmt.Fprintf(output, "Case #%d: %f\n", i+1, ret)
	}
}

func solve(D int64, N int64, K []int64, S []int64) float64 {
	// fmt.Println(R, C, L)

	var time float64 = float64(D-K[0]) / float64(S[0])

	for i := int64(0); i < N; i++ {
		this := float64(D-K[i]) / float64(S[i])
		if time < this {
			time = this
		}
	}

	return float64(D) / time
}

func readInt() int64 {
	var i int64
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
