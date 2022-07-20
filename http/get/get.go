package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	for {
		bs := make([]byte, 1024)
		n, err := response.Body.Read(bs)
		fmt.Println(string(bs[:n]))
		if err != nil || n == 0 {
			break
		}
	}

	// сокращенный вариант вывода на консоль
	// io.Copy(os.Stdout, response.Body)
}
