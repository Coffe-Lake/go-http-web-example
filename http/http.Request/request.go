package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest(
		"GET", "https://google.com", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // Добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // Добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
