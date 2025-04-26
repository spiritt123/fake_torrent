package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	GetPort  = "GET_PORT"
	SendPort = "SEND_PORT"
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
	case message == GetPort:
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

func startInteractiveMode(listenPort, peerAddress string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Клиент готов к вводу команд")
	fmt.Println("Доступные команды:")
	fmt.Printf("- %s <адрес>: получить порт указанного клиента\n", GetPort)
	fmt.Printf("- %s <адрес>: отправить свой порт указанному клиенту\n", SendPort)
	fmt.Printf("Текущий порт: %s\n", listenPort)
	if peerAddress != "" {
		fmt.Printf("Пир по умолчанию: %s\n", peerAddress)
	}

	for {
		fmt.Print("Введите команду: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]
		var target string
		if len(parts) == 2 {
			target = parts[1]
		} else if peerAddress != "" {
			target = peerAddress
		} else {
			fmt.Println("Не указан адрес клиента")
			continue
		}
		// add swih TODO add getPort sendPort
		switch cmd {
		// GetPort - value; getPort - function
		case GetPort:
			getPort(target)
		case SendPort:
			sendPort(target, listenPort)
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func getPort(target string) {
	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("Ошибка подключения к %s: %v\n", target, err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(GetPort + "\n"))
	if err != nil {
		fmt.Printf("Ошибка отправки GET_PORT: %v\n", err)
		return
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка чтения ответа: %v\n", err)
		return
	}

	fmt.Printf("Порт клиента %s: %s\n", target, strings.TrimSpace(response))
}

func sendPort(target, port string) {
	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("Ошибка подключения к %s: %v\n", target, err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(SendPort + " " + port + "\n"))
	if err != nil {
		fmt.Printf("Ошибка отправки SEND_PORT: %v\n", err)
		return
	}

	fmt.Printf("Порт %s успешно отправлен клиенту %s\n", port, target)
}
