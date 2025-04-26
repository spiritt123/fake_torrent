package main

import (
	"fmt"
	"os"
	"strings"
	"net"
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

func startServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера на порту %s: %v\n", port, err)
		return
	}
	defer listener.Close()
	fmt.Printf("Сервер слушает на порту %s\n", port)

	for {
		conn, err := listener.Accept();
		if err := nil {
			fmt.Printf("Ощибка при принятии соединения")
		}
		go handleConnection(conn, err)
	}
}