package main

import (
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("ok")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS user! Your config: %+v", r.TLS)
	})
	log.Fatal(http.Serve(autocert.NewListener(os.Args[1]), mux))
}
