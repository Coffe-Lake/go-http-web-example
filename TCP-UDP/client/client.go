package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Запуск...")
	for {
		var color string
		fmt.Println("Введите цвет: ")
		_, err := fmt.Scanln(&color)
		if err != nil {
			fmt.Println("Некорректный ввод/пустая строка!")
			continue
		}
		if color == "0" {
			fmt.Println("Выход...")
			break
		}
		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(color)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Перевод: ")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
		fmt.Println()
	}
}
