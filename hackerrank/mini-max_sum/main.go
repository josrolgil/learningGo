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

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int64)(int64, int64){
	const arrSize = 5
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var min int64 = 0
	var max int64 = 0

	for i:= 0; i<arrSize; i++{
		if i==0 {
			min += arr[i]
		}else if i==arrSize-1{
			max += arr[i]
		}else{
			min+=arr[i]
			max+=arr[i]
		}
	}
	fmt.Println(min, max)
	return min, max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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

