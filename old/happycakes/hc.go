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
		challenge := scanner.Text()
		res := calculate(challenge)
		fmt.Print("Case #", counter, ": ", res)
		fmt.Println()
		counter += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func calculate(state string) string {
	res := 0
	lastChar := '+'

	for _, c := range(Reverse(state)) {
		if c != lastChar {
			res += 1
			lastChar = c
		}
	}

	return strconv.Itoa(res)
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
