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
	R, C, H, V int
	Chips      [][]bool
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.R, c.C, c.H, c.V = readInt(), readInt(), readInt(), readInt()
		c.Chips = make([][]bool, c.R)
		for i := 0; i < c.R; i++ {
			c.Chips[i] = make([]bool, c.C)
			for j := 0; j < c.C; j++ {
				c.Chips[i][j] = readChar() == '@'
			}
			readChar()
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		fmt.Fprintf(os.Stderr, "Case #%v\n", c)
		res := solve(c.R, c.C, c.H, c.V, c.Chips)
		fmt.Printf("Case #%d: %s\n", i+1, res)
	}
}

func solve(R, C, H, V int, Chips [][]bool) string {
	return "IMPOSSIBLE"
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
