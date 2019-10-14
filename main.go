package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/berinle/rabbitmq-gopher/receive"
	"github.com/berinle/rabbitmq-gopher/send"
)

var port = 8888

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func sendHandler(w http.ResponseWriter, r *http.Request) {
	send.Send()
	fmt.Fprintf(w, "Message sent.")
}
func receiveHandler(w http.ResponseWriter, r *http.Request) {
	receive.Receive()
}

func main() {

	if p := os.Getenv("PORT"); p != "" {
		i, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal("PORT Conversion: ", err)
		}
		port = i
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/receive", receiveHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
