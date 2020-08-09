package main

import (
	"fmt"
	"io"
	"os"
)

func closeFP(pFP *os.File) {
	if pFP != nil {
		pFP.Close()
	}
}

func main() {
	for _, fpath := range os.Args[1:] {
		fmt.Println("Read content from ", fpath)
		fp, err := os.Open(fpath)
		if err != nil {
			fmt.Println("Fail to open file=", fpath)
			os.Exit(1)
		}
		defer closeFP(fp)
		io.Copy(os.Stdout, fp)
	}
}
