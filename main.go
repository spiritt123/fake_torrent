package main

import (
	"bufio"
	"fmt"
	"net"
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
	go startServer(listenPort)
	// TODO add startInteractive mode like menu
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
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Ошибка при принятии соединения: %v\n", err)
			continue
		}
		// to do handleConnection
		go handleConnection(conn, port)
	}
}

func handleConnection(conn net.Conn, serverPort string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка чтения команды: %v\n", err)
		return
	}

	message = strings.TrimSpace(message)
	switch {
	case message == cmdGetPort:
		_, err := conn.Write([]byte(serverPort + "\n"))
		if err != nil {
			fmt.Printf("Ошибка отправки порта: %v\n", err)
		}
	case strings.HasPrefix(message, SendPort):
		parts := strings.SplitN(message, " ", 2)
		if len(parts) < 2 {
			fmt.Println("Неверный формат команды SEND_PORT")
			return
		}
		fmt.Printf("Получен порт от клиента: %s\n", parts[1])
	default:
		fmt.Println("Неизвестная команда:", message)
	}
}
