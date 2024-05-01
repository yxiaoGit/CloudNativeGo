package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/throttle"
)

var throttled = throttle.Throttle(getHostname, 0, 0, time.Second)

func getHostname(ctx context.Context) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	return os.Hostname()
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func throttledHandler(w http.ResponseWriter, r *http.Request) {
	////r.RemoteAddr keeps changing
	ok, hostname, err := throttled(r.Context(), "localHost")
	log.Println(ok)
	log.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hostname))
}

func main() {
	//router := http.NewRouter()
	//mux := http.NewServeMux()
	http.HandleFunc("/hello", throttledHandler)

	//mux.HandleFunc("/hostn", throttledHandler)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

//curl http://localhost:3333/hello
