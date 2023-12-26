package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("a ...any")
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
