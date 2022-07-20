package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	httpRequest := "GET / HTTP/1.1\n" + "Host: google.com\n\n"
	conn, err := net.Dial("tcp", "google.com:80")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}
	// Создаем файл для записи
	newFile, err := os.Create("getres.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer newFile.Close()

	io.Copy(newFile, conn)
	fmt.Println("Done")
}
