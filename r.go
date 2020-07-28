package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:1324/arts/usr/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
