package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	cmdGetPort = "GET_PORT"
	SendPort   = "SEND_PORT"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("client <портПрослушивания> [файлИлиАдресПира]")
		return
	}

	listenPort := os.Args[1]
	var peerAddress string

	if len(os.Args) >= 3 {
		peerArg := os.Args[2]
		if _, err := os.Stat(peerArg); err == nil {
			content, err := os.ReadFile(peerArg)
			if err != nil {
				fmt.Printf("Ошибка чтения файла %s: %v\n", peerArg, err)
				return
			}
			peerAddress = strings.TrimSpace(string(content))
		} else {
			peerAddress = peerArg
		}
	}
	// TODO Add func relize startServer StartInteractive
	go startServer(listenPort)
	startInteractiveMode(listenPort, peerAddress)
}
