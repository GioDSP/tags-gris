package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

//le parametros na url e imprime no servidor
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//faz scan das portas de um endereco e ve se estao abertas ou fechadas
func scan(host string, port int, wg *sync.WaitGroup) error {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		log.Printf("\tPort %d: closed", port)
		return err
	}
	conn.Close()
	log.Printf("\tPort %d: open", port)
	return nil
}

func main() {
	//cria um servidor http na porta 80
	server := &http.Server{
		Addr:         "127.0.0.1:80",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second,
	}

	//var wg sync.WaitGroup
	//host := "localhost"

	log.Printf("Running web server on: http://%s\n", server.Addr)

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":80", nil)

	/*for port := 8070; port <= 8090; port++ {
		wg.Add(1)
		go scan(host, port, &wg)
	}

	log.Printf("Before Wait")
	wg.Wait()
	log.Printf("After Wait")
	*/
}
