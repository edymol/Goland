package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// the shorter way to get the response
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// this is the long way to get the response
	//defer resp.Body.Close() // Close the response body when done reading
	//
	//bs, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error reading response body:", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(string(bs))
}

// This is a custom implementation for Write function to accommodate the Copy interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("The number of bytes is: ", len(bs))

	return len(bs), nil
}
