package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	hostname string
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s %s %s\n", hostname, r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "Hostname: %s\n", hostname)
}

func HeaderHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s %s %s\n", hostname, r.RemoteAddr, r.Method, r.URL)
	header := r.URL.Query().Get("name")
	if header != "" {
		fmt.Fprintf(w, "Header %s value: %s\n", header, r.Header.Get(header))
		return
	}
	b, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, "Headers: %s\n", b)
}

func parseListen() string {
	var listenAddr string
	flag.StringVar(&listenAddr, "listen", "", "Server listen address")
	flag.Parse()
	if listenAddr != "" {
		return listenAddr
	}
	listenAddr, _ = os.LookupEnv("WHOAMI_LISTEN")
	if listenAddr != "" {
		return listenAddr
	}
	return "0.0.0.0:80"
}

func main() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	listenAddr := parseListen()

	log.Printf("Http server running on %s \n", listenAddr)
	http.HandleFunc("/", IndexHandle)
	http.HandleFunc("/headers", HeaderHandle)
	if err = http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalln(err)
	}
}
