package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
	"strings"
	"math"
	// "sort"
)

type InputData struct {
	A string
	B string
}

// Helpers
func read(input string) (int, []InputData) {
	ret := []InputData{}

	file, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())
	// fmt.Println("Total:", total)
	counter := 1
    for counter <= total {
		data := InputData{"", ""}

		scanner.Scan()
		readRow := strings.Split(scanner.Text(), " ")

		data.A = readRow[0]
		data.B = readRow[1]

		ret = append(ret, data)
		counter += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return total, ret
}

func main() {
	total, inputData := read(os.Args[1])

	i := 0
    for i < total {
		res := calculate(inputData[i])
		fmt.Print("Case #", (i+1), ": ", res)
		fmt.Println()
		i += 1
    }
}

func calculate(in InputData) string {
	out := InputData{"", ""}

	diff := getDiff(in.A, in.B)

	for i := 0; i < len(in.A); i++ {
		fmt.Println("I:", i)

		pow := int64(math.Pow10(len(in.A) - i - 1))

		if in.A[i] == '?' {
			choice := -1

			if in.B[i] == '?' {
				choice = 0
			} else {
				choice, _ = strconv.Atoi(string(in.B[i]))
			}

			firstChoice := choice
			firstDiff := diff

			fmt.Println("Try1", choice)
			newDiff := firstDiff + int64(choice) * pow
			fmt.Println(newDiff, "vs", diff)
			if absSmaller(newDiff, diff) {
				diff = newDiff
			}

			secondChoice := firstChoice - 1
			if secondChoice < 0 {
				secondChoice = 9
			}
			fmt.Println("Try2", secondChoice)
			newDiff = firstDiff + int64(secondChoice) * pow
			fmt.Println(newDiff, "vs", diff)
			if absSmaller(newDiff, diff) {
				diff = newDiff
				choice = secondChoice
			}

			thirdChoice := firstChoice + 1
			if thirdChoice >= 10 {
				thirdChoice = 0
			}
			fmt.Println("Try3", thirdChoice)
			newDiff = firstDiff + int64(thirdChoice) * pow
			fmt.Println(newDiff, "vs", diff)
			if absSmaller(newDiff, diff) {
				diff = newDiff
				choice = thirdChoice
			}

			fmt.Println("Chose", strconv.Itoa(choice))
			out.A = out.A + strconv.Itoa(choice)
		} else {
			fmt.Println("Kept", string(in.A[i]))
			out.A = out.A + string(in.A[i])
		}

		if in.B[i] == '?' {
			choice, _ := strconv.Atoi(string(in.A[i]))

						firstChoice := choice
			firstDiff := diff

			fmt.Println("Try1", choice)
			newDiff := firstDiff - int64(choice) * pow
			fmt.Println(newDiff, "vs", diff)
			if absSmaller(newDiff, diff) {
				diff = newDiff
			}

			secondChoice := firstChoice - 1
			if secondChoice < 0 {
				secondChoice = 9
			}
			fmt.Println("Try2", secondChoice)
			newDiff = firstDiff - int64(secondChoice) * pow
			fmt.Println(newDiff, "vs", diff)
			if absSmaller(newDiff, diff) {
				diff = newDiff
				choice = secondChoice
			}

			thirdChoice := firstChoice + 1
			if thirdChoice >= 10 {
				thirdChoice = 0
			}
			fmt.Println("Try3", thirdChoice)
			newDiff = firstDiff - int64(thirdChoice) * pow
			fmt.Println(newDiff, "vs", diff)
			if absSmaller(newDiff, diff) {
				diff = newDiff
				choice = thirdChoice
			}

			fmt.Println("Chose", strconv.Itoa(choice))

			out.B = out.B + strconv.Itoa(choice)
		} else {
			out.B = out.B + string(in.B[i])
		}
	}

	return out.A + " " + out.B
}

func getDiff(a string, b string) int64 {
	aint, _ := strconv.ParseInt(
		strings.Replace(a, "?", "0", -1), 10, 64)
	bint, _ := strconv.ParseInt(
		strings.Replace(b, "?", "0", -1), 10, 64)

	return aint - bint
}

func absSmaller(n int64, m int64) bool {
	if n < 0 {
		n = -n
	}
	if m < 0 {
		m = -m
	}

	return n < m
}
