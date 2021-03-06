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
		fmt.Println("Case", i+1, ":", c)
		s, _, _ := solve(c.Hd, c.Hd, c.Ad, c.Hk, c.Ak, c.B, c.D, 0, 0, nil)
		if s == -1 {
			// fmt.Println("IMPOSSIBLE", err, string(track))
			fmt.Fprintf(output, "Case #%d: IMPOSSIBLE\n", i+1)
		} else {
			// fmt.Println("OK", s, err, string(track))
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
	D int, counter int, state int, track []rune) (int, error, []rune) {
	// fmt.Println(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state, string(track))

	if counter > 10000 {
		return -1, errors.New("Over counter"), track
	}

	if Hk <= 0 {
		return counter, errors.New("Knight killed, counter: " + strconv.Itoa(counter)), track
	}

	if counter > 0 {
		Hd -= Ak
	}

	if Hd <= 0 {
		return -1, errors.New("Dragon killed"), track
	}

	// Special case 1: Knight can be killed
	if Ad >= Hk {
		// Attack
		return solve(HdInitial, Hd, Ad, Hk-Ad, Ak, B, D, counter+1, state, append(track, 'a'))
	}

	// Special case 2: Dragon must heal
	if Ak >= Hd {
		// Cure
		return solve(HdInitial, HdInitial, Ad, Hk, Ak, B, D, counter+1, state, append(track, 'c'))
	}

	switch state {
	case 0: // Debuff or move on
		if D <= 0 || Ak <= 0 {
			// Next state
			return solve(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state+1, track)
		}
		// Debuff
		d, errd, td := solve(HdInitial, Hd, Ad, Hk, Ak-D, B, D, counter+1, state, append(track, 'd'))
		// Next state
		n, errn, tn := solve(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state+1, track)
		if d == -1 || (n != -1 && n < d) {
			return n, errn, tn
		}
		return d, errd, td
	case 1: // Buff or move on
		if B <= 0 {
			// Next state
			return solve(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state+1, track)
		}
		// Buff
		b, errb, tb := solve(HdInitial, Hd, Ad+B, Hk, Ak, B, D, counter+1, state, append(track, 'b'))
		// Next state
		n, errn, tn := solve(HdInitial, Hd, Ad, Hk, Ak, B, D, counter, state+1, track)
		if b == -1 || (n != -1 && n < b) {
			return n, errn, tn
		}
		return b, errb, tb
	case 2: // Attack always
		// Attack
		return solve(HdInitial, Hd, Ad, Hk-Ad, Ak, B, D, counter+1, state, append(track, 'a'))
	}

	panic("No")
	return -1, errors.New("NO"), track
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
