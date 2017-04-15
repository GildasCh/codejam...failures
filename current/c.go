package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
		fmt.Println("Case", i, ":", c)
		s, err := solve(c.Hd, c.Hd, c.Ad, c.Hk, c.Ak, c.B, c.D, 0, 0)
		if s == -1 {
			fmt.Println("IMPOSSIBLE", err)
			fmt.Fprintf(output, "Case #%d: IMPOSSIBLE\n", i+1)
		} else {
			fmt.Println("OK", s, err)
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
	D int, counter int, state int) (int, error) {
	// fmt.Println(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state)

	if counter > 1000 {
		return -1, errors.New("Over counter")
	}

	if Hk <= 0 {
		return counter, errors.New("Knight killed, counter: " + strconv.Itoa(counter))
	}

	if counter > 0 {
		Hd -= Ak
	}

	if Hd <= 0 {
		return -1, errors.New("Dragon killed")
	}

	// Special case 1: Knight can be killed
	if Ad >= Hk {
		// Attack
		return solve(HdInitial, Hd, Ad, Hk-Ad, Ak, B, D, counter+1, state)
	}

	// Special case 2: Dragon must heal
	if Ak >= Hd {
		// Cure
		return solve(HdInitial, HdInitial, Ad, Hk, Ak, B, D, counter+1, state)
	}

	switch state {
	case 0: // Debuff or move on
		if Ak <= 0 {
			Ak = 0
			// Next state
			return solve(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state+1)
		}
		// Debuff
		d, err := solve(HdInitial, Hd, Ad, Hk, Ak-D, B, D, counter+1, state)
		// Next state
		n, err := solve(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state+1)
		return min(d, n), err
	case 1: // Buff or move on
		// Buff
		b, err := solve(HdInitial, Hd, Ad+B, Hk, Ak, B, D, counter+1, state)
		// Next state
		n, err := solve(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state+1)
		return min(b, n), err
	case 2: // Attack always
		// Attack
		return solve(HdInitial, Hd, Ad, Hk-Ad, Ak, B, D, counter+1, state)
	}

	panic("No")
	return -1, errors.New("NO")
}

func min(a, b int) int {
	if a == -1 || (b != -1 && b < a) {
		return b
	}
	return a
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
