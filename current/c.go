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
	N uint64
	K uint64
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
		c.N = readUint64()
		c.K = readUint64()
		C = append(C, c)
	}

	for i, c := range C {
		max, min := solve(c.N, c.K)
		fmt.Fprintf(output, "Case #%d: %d %d\n", i+1, max, min)
	}
}

func solve(N uint64, K uint64) (uint64, uint64) {
	numberOfPeopleSetup := K - 1
	numberOfSpaces := K
	totalFreeSpace := N - numberOfPeopleSetup

	lastEquitableShare := floorTwoPower(numberOfSpaces)

	smallSpaces := (N - lastEquitableShare + 1) / lastEquitableShare
	excess := (N - lastEquitableShare + 1) % lastEquitableShare

	chosenSpace := smallSpaces
	if K-lastEquitableShare < excess {
		chosenSpace++
	}

	min := (chosenSpace - 1) / 2
	max := min
	if (chosenSpace-1)%2 > 0 {
		max++
	}

	fmt.Println("N", N)
	fmt.Println("K", K)
	fmt.Println("lastEquitableShare", lastEquitableShare)
	fmt.Println("totalFreeSpace", totalFreeSpace)
	fmt.Println("smallSpaces", smallSpaces)
	fmt.Println("excess", excess)
	fmt.Println("chosenSpace", chosenSpace)
	fmt.Println("min", min)
	fmt.Println("max", max)
	return max, min
}

func floorTwoPower(spaces uint64) uint64 {
	var ret uint64 = 1

	for {
		if ret > spaces {
			return ret / 2
		}
		ret = ret * 2
	}

	panic("Not supposed to be here")
	return 0
}

func readInt() int {
	var i int
	fmt.Fscanf(input, "%d", &i)
	return i
}

func readUint64() uint64 {
	var i uint64
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
