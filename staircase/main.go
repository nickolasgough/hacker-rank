package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the staircase function below.
// Found at https://www.hackerrank.com/challenges/staircase/problem.
func staircase(n int32) {
	s := n - 1
	for s >= 0 {
		var r string
		for i := int32(1); i <= s; i++ {
			r += " "
		}
		for i := int32(1); i <= (n - s); i++ {
			r += "#"
		}
		fmt.Println(r)
		s--
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
