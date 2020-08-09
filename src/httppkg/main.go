package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	pResp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Something wrong:", err)
		os.Exit(1)
	}

	// Show struct Response content
	// fmt.Println(p_resp)

	/* bsize := 1024 * 1024 // 1 Mbyte
	bs := make([]byte, bsize)
	readByteSize, err := pResp.Body.Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println("Something wrong: ", err)
		os.Exit(1)
	}
	fmt.Println("Read ", readByteSize, " bytes")
	fmt.Println(string(bs[:readByteSize])) */

	// Use Writer interface to show content
	// io.Copy(os.Stdout, pResp.Body)
	lw := logWriter{}
	io.Copy(lw, pResp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
