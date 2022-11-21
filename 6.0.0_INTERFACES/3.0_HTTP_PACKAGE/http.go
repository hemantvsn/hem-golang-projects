package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//simpleHttpGet()
	//advancedHttpGet()
	httpClientWithCustomStringWriter()
}

// Which implements write interface
type StringWriter struct {
	value string
}

// Implements writer
func (sw StringWriter) Write(bytes []byte) (n int, err error) {
	sw.value = string(bytes)
	fmt.Println("The sw =", sw)
	return len(sw.value), nil
}

func httpClientWithCustomStringWriter() {
	var resp, err = http.Get("https://dummyjson.com/products/1")
	if err != nil {
		fmt.Println("Error occured while fetching data", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Got response from API but its not 200", resp)
		os.Exit(1)
	}

	io.Copy(StringWriter{}, resp.Body)

}

// Uses writer interface
func advancedHttpGet() {
	var resp, err = http.Get("https://dummyjson.com/products/1")
	if err != nil {
		fmt.Println("Error occured while fetching data", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Got response from API but its not 200", resp)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}

// Basic implementation
func simpleHttpGet() {

	var resp, err = http.Get("https://dummyjson.com/products/1")
	if err != nil {
		fmt.Println("Error occured while fetching data", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Got response from API but its not 200", resp)
		os.Exit(1)
	}

	// var bytes = []byte{} => DOESN'T WORK
	// This is because READ function will read data into byteSlice until its full
	// But it wont resize byteSlice
	// If we pass empty slice, read function will see it as full and wont read data into it
	// Hence pass slice with higher capacity
	var bytes = make([]byte, 1000)
	var num, err1 = resp.Body.Read(bytes)

	fmt.Println(num, err1)
	fmt.Println(string(bytes))
}
