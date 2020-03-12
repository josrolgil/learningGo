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
func staircase(n int32) {
	var aux int = 1
	for i:=0; i<int(n); i++{
		generateString(int(n)-aux-i,i+aux)
		fmt.Print("\n")
	}

}

func generateString(spacesNumber int, symbolsNumber int){
	for i := 0; i < spacesNumber; i++ {
		fmt.Print(" ")
	}
	for i := 0; i < symbolsNumber; i++ {
		fmt.Print("#")
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
