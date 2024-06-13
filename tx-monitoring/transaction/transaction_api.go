package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %v not allowed", r.Method)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error reading request body: %v\n", err)
		return
	}

	var tx Transaction
	err = json.Unmarshal(body, &tx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("error parse JSON: %v\n", err)
		return
	}

	hash, err := tx.Broadcast()
	if err != nil {
		fmt.Println("error transaction broadcast:", err)
		return
	}
	resData := TxHashResponse{
		TxHash: hash,
	}
	jsonData, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("error marshal data to json:", err)
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("error writing the response:", err)
	}
}

func GetTransactionStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %v not allowed", r.Method)
		return
	}

	hash := r.URL.Path[len("/get-tx-status/"):]

	retry := 2
	monitorInterval := 5
	status, err := Monitor(hash, retry, monitorInterval)
	if err != nil {
		fmt.Println("error transaction monitoring:", err)
		return
	}

	resData := TxStatusResponse{
		TxStatus: status,
	}
	jsonData, err := json.Marshal(resData)
	if err != nil {
		fmt.Println("error marshal data to json:", err)
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("error writing the response:", err)
	}
}
