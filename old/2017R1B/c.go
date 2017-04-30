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
	N int
	Q int
	E []float64
	S []float64
	D []map[int]float64
	U []int
	V []int
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
		c.N = readInt()
		c.Q = readInt()

		c.E = make([]float64, c.N)
		c.S = make([]float64, c.N)
		for j := 0; j < c.N; j++ {
			c.E[j] = readFloat()
			c.S[j] = readFloat()
		}

		c.D = make([]map[int]float64, c.N)
		for j := 0; j < c.N; j++ {
			c.D[j] = make(map[int]float64)
			for k := 0; k < c.N; k++ {
				read := readInt()
				if read == -1 {
					continue
				}
				c.D[j][k] = float64(read)
			}
		}

		c.U = make([]int, c.Q)
		c.V = make([]int, c.Q)
		for j := 0; j < c.Q; j++ {
			c.U[j] = readInt()
			c.V[j] = readInt()
		}

		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		res := solve(c.N, c.Q, c.E, c.S, c.D, c.U, c.V)
		fmt.Fprintf(output, "Case #%d: %f\n", i+1, res)
	}
}

func solve(N, Q int, E, S []float64, D []map[int]float64, U, V []int) float64 {
	fmt.Println(N, Q, E, S, D, U, V)
	return 1.0
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
