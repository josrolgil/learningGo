package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	hourSystem :=s[8:10]
	if hourSystem == "AM"{
		return s[0:8]
	}


	var hourString string
	if(s[0]==0){
		hourString=s[1:2]
	}else{
		hourString=s[0:2]
	}
	hour, _ := strconv.Atoi(hourString)
	hour +=12 //Converting from 12 to 24
	newTime := strconv.Itoa(hour)+":"+s[3:5]+":"+s[6:8]
	return newTime

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
