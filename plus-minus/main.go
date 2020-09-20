package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
// Found at https://www.hackerrank.com/challenges/plus-minus/problem.
func plusMinus(numbers []int32) {
	var nPositive, nNegative, nZero float32
	for _, number := range numbers {
		if number > 0 {
			nPositive++
		} else if number < 0 {
			nNegative++
		} else {
			nZero++
		}
	}
	nNumbers := float32(len(numbers))
	fmt.Printf("%.6f\n", nPositive / nNumbers)
	fmt.Printf("%.6f\n", nNegative / nNumbers)
	fmt.Printf("%.6f", nZero / nNumbers)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
