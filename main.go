package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	hostname   string
	listenAddr string
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s %s %s\n", hostname, r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "Hostname: %s\n", hostname)
}

func main() {
	flag.StringVar(&listenAddr, "listen", ":80", "Server listen address")
	flag.Parse()
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Http server running on %s \n", listenAddr)
	http.HandleFunc("/", IndexHandle)
	if err = http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalln(err)
	}
}
