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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	var rotatedString string
	k %= 26
	for _, alphabet := range s {
		if alphabet >= 'a' && alphabet <= 'z' {
			if alphabet = (alphabet + k); alphabet > 'z' {
				alphabet = alphabet - 'z' + 'a' - 1
			}
		} else if alphabet >= 'A' && alphabet <= 'Z' {
			if alphabet = (alphabet + k); alphabet > 'Z' {
				alphabet = alphabet - 'Z' + 'A' - 1
			}
		}
		rotatedString += string(alphabet)
	}
	return rotatedString
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

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
