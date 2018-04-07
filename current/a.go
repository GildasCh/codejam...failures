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
	T = readInt()

	for i := 0; i < T; i++ {
		c := Case{}
		c.D = readInt()
		c.P = readString()
		C = append(C, c)
	}

	fmt.Fprintf(os.Stderr, "Solving...\n")

	for i, c := range C {
		fmt.Fprintf(os.Stderr, "Case #%d\n", i+1)
		res, ok := solve(c.D, c.P)
		if !ok {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
		} else {
			fmt.Printf("Case #%d: %d\n", i+1, res)
		}
	}
}

func solve(D int, P string) (int, bool) {
	fmt.Fprintf(os.Stderr, "solving %d %q\n", D, P)
	fmt.Fprintf(os.Stderr, "power: %q\n", power(P))

	if D < numberOfS(P) {
		return -1, false
	}

	n := 0

	for power(P) > D {
		P = swap(P)
		n++
	}

	return n, true
}

func swap(P string) string {
	program := []byte(P)

	for i := len(program) - 2; i >= 0; i-- {
		if program[i] == 'C' && program[i+1] == 'S' {
			program[i], program[i+1] = program[i+1], program[i]
			return string(program)
		}
	}

	return "PROUT"
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
	fmt.Fscanf(os.Stdin, "%d", &i)
	return i
}

func readString() string {
	var str string
	fmt.Fscanf(os.Stdin, "%s", &str)
	return str
}
