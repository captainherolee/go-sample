package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var image []byte

func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

func handleerHtml(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		fmt.Println("push /image")
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	http.HandleFunc("/", handleHtml)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18433", "cert.pem", "key.pem", nil)
	fmt.Println(err)
}
