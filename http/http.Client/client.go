package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	duration := time.Since(start)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())
}
