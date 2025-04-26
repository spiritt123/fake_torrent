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

	for {
		displayClearScreen()
		displayHeaderOnScreen(listenPort)
		displayMenu()
		if peerAddress != "" {
			fmt.Printf("\nТекущий пир %s\n", peerAddress)
		}

		fmt.Print("\nВыберите действие (введите номер команды или команду): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch {
		case input == "1", string.EqualFold(input, GetPort):
			handleGetPort(reader, listenPort, PeerAddres)
		case input == "2", strings.EqualFold(input, SendPort):
			handleSendPort(reader, listenPort, peerAddress)
		case input == "3", strings.EqualFold(input, "SET_PEER"):
			peerAddress = handleSetPeer(reader)
		case input == "4", strings.EqualFold(input, "HELP"):
			continue
		case input == "5", strings.EqualFold(input, "EXIT"), strings.EqualFold(input, "QUIT"):
			fmt.Println("\nЗавершение работы клиента...")
			return
		default:
			fmt.Println("\nНеизвестная команда. Введите HELP для списка команд.")
			promptContinue(reader)
		}

	}

}

func displayClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func displayHeaderOnScreen(port string) {
	fmt.Println("=== Файловый клиент ===")
	fmt.Printf("Ваш порт: %s\n", port)
	fmt.Println("=======================")
}
func displayMenu() {
	fmt.Println("\nМеню:")
	fmt.Println("1. GET_PORT  - Получить порт другого клиента")
	fmt.Println("2. SEND_PORT - Отправить свой порт другому клиенту")
	fmt.Println("3. SET_PEER  - Изменить адрес пира по умолчанию")
	fmt.Println("4. HELP      - Показать это меню")
	fmt.Println("5. EXIT      - Выйти из программы")
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
