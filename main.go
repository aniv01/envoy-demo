package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "hello! i'm Mr.%v", hostname)
}

func admin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is an admin page!")
}

func main() {

	port := flag.String("port", "8080", "port number")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/admin", admin)

	log.Println("starting server on port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, mux))
}
