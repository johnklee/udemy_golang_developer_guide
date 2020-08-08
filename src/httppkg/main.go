package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	pResp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Something wrong:", err)
		os.Exit(1)
	}

	// Show struct Response content
	// fmt.Println(p_resp)
}
