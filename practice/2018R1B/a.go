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
	N, L int
	C    []int
}

func main() {
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.N, c.L = readInt(), readInt()
		c.C = make([]int, c.L)
		for i := 0; i < c.L; i++ {
			c.C[i] = readInt()
		}
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		fmt.Fprintf(os.Stderr, "Case #%v\n", c)
		res := solve(c.N, c.L, c.C)
		fmt.Printf("Case #%d: %d\n", i+1, res)
	}
}

func solve(N, L int, C []int) int {
	current := goods(N, C)
	voted := sum(C)

	toAdd := 1
	for {
	start:
		if N < voted+toAdd {
			C[0] += toAdd - 1
			fmt.Fprintf(os.Stderr, "final conf: %v (score: %d)\n", C, score(N, C))
			return score(N, C)
		}

		for i := 0; i < L; i++ {
			C[i] += toAdd
			if goods(N, C) > current {
				fmt.Fprintf(os.Stderr, "found better: %v (score: %d)\n", C, goods(N, C))
				current = goods(N, C)
				voted++
				toAdd = 1
				goto start
			}
			C[i] -= toAdd
		}

		C2 := append(C, toAdd)
		if goods(N, C2) > current {
			C = C2
			fmt.Fprintf(os.Stderr, "found better: %v (score: %d)\n", C, goods(N, C))
			current = goods(N, C2)
			voted++
			toAdd = 1
			goto start
		}

		toAdd++
	}

	return -1
}

func sum(C []int) int {
	sum := 0
	for _, c := range C {
		sum += c
	}
	return sum
}

func score(N int, C []int) int {
	sum := 0
	for _, c := range C {
		d := 100 * float64(c) / float64(N)
		if d >= float64(float64(int(d))+0.5) {
			sum += int(d) + 1
		} else {
			sum += int(d)
		}
	}

	return sum
}

func goods(N int, C []int) int {
	sum := 0
	for _, c := range C {
		d := 100 * float64(c) / float64(N)
		if d >= float64(float64(int(d))+0.5) {
			sum += 1
		}
	}

	return sum
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
