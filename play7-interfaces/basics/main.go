package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type out struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(resp)
	// 1. we define a byte array and pass it to the the read fn
	output := make([]byte, 9999)
	resp.Body.Read(output)
	fmt.Println(string(output))

	// 2. we can also directly pass the stdout writer fn to the io.copy function
	io.Copy(os.Stdout, resp.Body)

	// 3. we write our own interface to capture the byte stream and print to sysout
	customWriter := out{}
	io.Copy(customWriter, resp.Body)
}

// implemeting the Write interface which is receives type out as the struct
func (out) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
