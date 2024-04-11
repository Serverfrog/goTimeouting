package main

import (
	"net/http"
	"time"
)

type timeOutHandler struct {
}

func (th timeOutHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(time.Second * 90)

	_, err := w.Write([]byte("you were waiting 90 seconds"))
	if err != nil {
		println(err)
	}
}

func main() {
	to := timeOutHandler{}
	println("Server starting. Timout is 90 Seconds")
	err := http.ListenAndServe(":8080", to)
	if err != nil {
		panic(err)
	}
}
