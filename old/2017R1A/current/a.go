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
	R int
	C int
	L [][]rune
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
		c.R = readInt()
		c.C = readInt()
		c.L = make([][]rune, c.R)
		for i := 0; i < c.R; i++ {
			// c.L[i] = make([]rune, c.C)
			line := readString()
			c.L[i] = []rune(line)
		}
		C = append(C, c)
	}

	fmt.Println("Solving...")

	for i, c := range C {
		solve(c.R, c.C, c.L)
		fmt.Fprintf(output, "Case #%d:\n", i+1)
		for _, l := range c.L {
			fmt.Fprintf(output, "%s\n", string(l))
		}
	}
}

func solve(R int, C int, L [][]rune) int {
	// fmt.Println(R, C, L)

	for i := 0; i < R; i++ {
		first := firstNonEmpty(L[i])
		// fmt.Println("firstNonEmpty:", first)

		if first == -1 {
			if i > 0 {
				for j := 0; j < C; j++ {
					L[i][j] = L[i-1][j]
				}
			}
			continue
		}

		current := L[i][first]
		for j := 0; j < C; j++ {
			if L[i][j] == '?' {
				L[i][j] = current
			} else {
				current = L[i][j]
			}
		}
	}

	// Fix first lines
	first := firstNonEmptyRow(L)
	for i := 0; i < first; i++ {
		for j := 0; j < C; j++ {
			L[i][j] = L[first][j]
		}
	}

	return 1
}

func firstNonEmpty(line []rune) int {
	for i, r := range line {
		if r != '?' {
			return i
		}
	}
	return -1
}

func firstNonEmptyRow(L [][]rune) int {
	for i := 0; i < len(L); i++ {
		if L[i][0] != '?' {
			return i
		}
	}

	return -1
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
