package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
)

func main() {
	//fmt.Println(getDigitsFromInt(2847))
	// return

	file, err := os.Open(os.Args[1])
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
        // fmt.Println(scanner.Text())
		scanner.Scan()
		challenge, _ := strconv.Atoi(scanner.Text())
		res := calculate(challenge)
		fmt.Print("Case #", counter, ": ", res)
		fmt.Println()
		counter += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func calculate(n int) string {
	digits := make([]int, 0, 0)

	k := 1
	for true {
		// fmt.Println(k, ";", digits)
		digits = addDigits(digits, getDigitsFromInt(k*n))
		if len(digits) >= 10 {
			return strconv.Itoa(k*n)
		}
		if k*n == (k-1)*n {
			return "INSOMNIA"
		}
		k += 1
	}

	return "ERROR"
}

func getDigitsFromInt(number int) []int {
	res := make([]int, 0, 0)
	for number > 0 {
		res = appendIfMissing(res, number % 10)
		number = number / 10
	}
	return res
}

func addDigits(s1 []int, s2 []int) []int {
	for _, v := range s2 {
		s1 = appendIfMissing(s1, v)
	}
	return s1
}

func appendIfMissing(slice []int, i int) []int {
    for _, ele := range slice {
        if ele == i {
            return slice
        }
    }
    return append(slice, i)
}
