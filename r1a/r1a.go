package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
	// "strings"
	// "math"
)

type InputData struct {
	S string
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
        // fmt.Println(scanner.Text())
		scanner.Scan()
		data.S = scanner.Text()

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
	out := ""

	var first rune

	for _, c := range in.S {
		if out == "" {
			out = out + string(c)
			first = c
		} else {
			if c >= first {
				out = string(c) + out
				first = c
			} else {
				out = out + string(c)
			}
		}
	}

	return out
}
