package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// How to make an http request
func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//Declare byte slice with different syntax
	// make a slice of type byte
	//Make takes a type of slice (array) and the second element is the number of empy spaces
	bs := make([]byte, 99999)

	//Pass the byte to read function
	//The read function will grab the http info and put in the byte
	//resp body takes the byte, then grabs the html in the body and puts it inside the byte
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	//Different way to read, smaller
	//In the line bellow we use the writer interface
	//Takes data and sends some form of output
	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
