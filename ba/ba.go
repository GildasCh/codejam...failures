package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
	"strings"
	// "math"
	"sort"
)

type InputData struct {
	A string
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
		data := InputData{""}
		scanner.Scan()
		data.A = scanner.Text()

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
	numbers := getNumbers(in.A)

	sort.Ints(numbers)

	ret := ""
	for _, n := range(numbers) {
		ret += strconv.Itoa(n)
	}

	return ret
}

func getNumbers(num string) []int {
	ret := []int{}

	// First pass
	for (find(num, 'Z')) {
		num = remove(num, "ZERO")
		ret = append(ret, 0)
	}
	for (find(num, 'W')) {
		num = remove(num, "TWO")
		ret = append(ret, 2)
	}
	for (find(num, 'X')) {
		num = remove(num, "SIX")
		ret = append(ret, 6)
	}
	for (find(num, 'G')) {
		num = remove(num, "EIGHT")
		ret = append(ret, 8)
	}

	// Second Pass
	for (find(num, 'H')) {
		num = remove(num, "THREE")
		ret = append(ret, 3)
	}
	for (find(num, 'S')) {
		num = remove(num, "SEVEN")
		ret = append(ret, 7)
	}

	// Third pass
	for (find(num, 'R')) {
		num = remove(num, "FOUR")
		ret = append(ret, 4)
	}
	for (find(num, 'V')) {
		num = remove(num, "FIVE")
		ret = append(ret, 5)
	}

	// Fourth pass
	for (find(num, 'O')) {
		num = remove(num, "ONE")
		ret = append(ret, 1)
	}
	for (find(num, 'I')) {
		num = remove(num, "NINE")
		ret = append(ret, 9)
	}

	return ret
}

func find(in string, query rune) bool {
	for _, c := range(in) {
		if c == query {
			return true
		}
	}
	return false
}

func remove(in string, word string) string {
	ret := in

	for _, w := range(word) {
		ret = strings.Replace(ret, string(w), "", 1)
	}
	return ret
}
