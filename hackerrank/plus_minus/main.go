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
func plusMinus(arr []int32) {
	var counterPositive, counterNegative, counterZero int32= 0,0,0
	for _, element := range arr{
		if element==0 {
			counterZero++
		}else if element>0 {
			counterPositive++
		}else{
			counterNegative++
		}
	}

	printer(calculation(counterPositive, len(arr)), calculation(counterNegative, len(arr)), calculation(counterZero, len(arr)))
}

func calculation(counter int32, total int) float32{
	return float32(counter)/float32(total)
}

func printer(positiveFraction float32, negativeFraction float32, zeroFraction float32){
	fmt.Printf("%.6f\n",positiveFraction)
	fmt.Printf("%.6f\n",negativeFraction)
	fmt.Printf("%.6f\n",zeroFraction)

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
