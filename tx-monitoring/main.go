package main

import (
	"fmt"
	"net/http"
	"ryosantouchh/tx-monitoring/transaction"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %v not allowed", r.Method)
		return
	}

	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/get-tx-status/", transaction.GetTransactionStatusHandler)
	http.HandleFunc("/create-tx", transaction.CreateTransactionHandler)

	port := ":8080"
	fmt.Println("start server at port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("failed to start server at port:", port)
	}
}
