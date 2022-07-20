package main

import (
	"fmt"
	"net"
)

var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func main() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error", err)
			break
		}
		source := string(input[0:n])
		// на основании полученных данных получаем из словаря перевод
		target, ok := dict[source]
		if !ok { // если данные не найдены в словаре
			target = "undefined"
		}
		// выводим на консоль сервера диагностическую информацию
		fmt.Println(source, "-", target)
		// отправляем данные клиенту
		conn.Write([]byte(target))
	}
}
