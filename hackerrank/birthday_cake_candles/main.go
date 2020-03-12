package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(arr []int32) int32 {
	//Case array is one element
	if(len(arr) == 1){
		return 1
	}
	sortDescending(arr)

	if(arr[0]>arr[1]){
		return 1
	}else{
		var i int32=0
		for i=1; i<int32(len(arr)); i++{
			if arr[i-1]>arr[i] {
				return i
			}
		}
		return i
	}
}

func sortDescending(arr []int32){
	sort.SliceStable(arr, func(i,j int)bool { return arr[i] > arr[j]})
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := birthdayCakeCandles(ar)

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
