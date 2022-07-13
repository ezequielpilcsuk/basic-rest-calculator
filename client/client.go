package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	baseURL := "http://localhost:9091/"

	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Printf("erro: %d\n", err)
		return
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	// check errors
	fmt.Println(buf.String())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Available operations: add, sub, mul, div\nSelect the operation: ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSuffix(operation, "\n")

	fmt.Print("Enter the first value: ")
	value1, _ := reader.ReadString('\n')
	value1 = strings.TrimSuffix(value1, "\n")

	fmt.Print("Enter the second value: ")
	value2, _ := reader.ReadString('\n')
	value2 = strings.TrimSuffix(value2, "\n")

	fullRequest := fmt.Sprintf("%vcalculate?operation=%v&value1=%v&value2=%v",
		baseURL, operation, value1, value2)

	resp, err = http.Get(fullRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf = new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	// check errors
	fmt.Println(buf.String())

}
