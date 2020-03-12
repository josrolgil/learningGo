package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the birthday function below.
func birthday(s []int32, d int32, m int32) int32 {
	result := int32(0)

	if m > int32(len(s)) {
		return 0
	}

	if len(s) == 1 {
		if isValid(s, d) {
			return 1
		}
	}

	for i := int32(0); i < int32(len(s))-m+1; i++ {
		share := s[i : i+m]
		if isValid(share, d) {
			result++
		}
	}
	return result
}
func isValid(share []int32, d int32) bool {
	sum := calculateSum(share)
	if sum == d {
		return true
	}
	return false
}
func calculateSum(share []int32) int32 {
	sum := int32(0)
	for j := 0; j < len(share); j++ {
		sum += share[j]
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	dm := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	dTemp, err := strconv.ParseInt(dm[0], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(dm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := birthday(s, d, m)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
