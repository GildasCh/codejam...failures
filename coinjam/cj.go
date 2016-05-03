package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"fmt"
	"strings"
	"math"
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
		challenge := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(challenge[0])
		j, _ := strconv.Atoi(challenge[1])
		res := calculate(n, j)
		fmt.Print("Case #", counter, ":\n", res)
		fmt.Println()
		counter += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func calculate(n int, j int) string {
	answers := make([]string, 0, 0)

	candidate := 0
	for len(answers) < j {
		binCandidate := binaryString(candidate, n-2)
		if binCandidate == "" {
			return "NO SOLUTION"
		}

		binCandidate = "1" + binCandidate + "1"
		if ok, solution := isPrimeInAllBases(binCandidate); ok {
			fmt.Println("found solution", len(answers), ":", solution)
			answers = append(answers, binCandidate + " " + solution)
		}
		candidate += 1
	}

	return strings.Join(answers, "\n")
}

func binaryString(n int, length int) string {
	res := strconv.FormatInt(int64(n), 2)
	if len(res) > length {
		return ""
	}

	for len(res) < length {
		res = "0" + res
	}
	return res
}

func isPrimeInAllBases(bin string) (bool, string) {
	proofString := ""

	for base := 2; base <= 10; base++ {
		number := binStringToInt(bin, base)
		prime, divisor := isPrime(number)
		if prime {
			return false, ""
		}
		proofString += strconv.Itoa(divisor) + " "
	}
	return true, strings.Trim(proofString, " ")
}

func binStringToInt(bin string, base int) int64 {
	res := int64(0)
	for _, c := range(bin) {
		res *= int64(base)
		if c == '1' {
			res += 1
		}
	}
	return res
}

func isPrime(n int64) (bool, int) {
	sr := int(math.Sqrt(float64(n)))
	for i := 2; i < sr + 1; i++ {
		if n % int64(i) == 0 {
			return false, i
		}
	}
	return true, 0
}
