package main

import (
	"fmt"
	"os"
)

const (
	GetPort  = "need to port"
	SendPort = "need port"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("go 0")
		return
	}

}
