package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Transaction struct {
	Symbol    string `json:"symbol"`
	Price     int    `json:"price"`
	Timestamp int64  `json:"timestamp"`
	Hash      string
	Retry     int
}

type TxHashResponse struct {
	TxHash string `json:"tx_hash"`
}

type TxStatusResponse struct {
	TxStatus string `json:"tx_status"`
}

func (tx *Transaction) Broadcast() (string, error) {
	data, err := json.Marshal(tx)
	if err != nil {
		return "", err
	}

	res, err := http.Post("https://mock-node-wgqbnxruha-as.a.run.app/broadcast", "application/json", bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected/wrong status code: %v\n", res.StatusCode)
	}

	var response map[string]string
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}
	tx.Hash = response["tx_hash"]
	fmt.Printf("Created transaction : %v \nSymbol : %v \nPrice : %v \nTimestamp : %v \n \n", tx.Hash, tx.Symbol, tx.Price, tx.Timestamp)

	return tx.Hash, nil
}

func Monitor(hash string, retry int, interval int) (string, error) {
	fmt.Printf("start monitoring transaction: %v\n", hash)
	for i := 0; i < retry; i++ {
		res, err := http.Get(fmt.Sprintf("https://mock-node-wgqbnxruha-as.a.run.app/check/%v", hash))
		if err != nil {
			return "", err
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

		var response map[string]string
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return "", err
		}

		status := response["tx_status"]
		switch status {
		case "CONFIRMED":
			fmt.Printf("transaction: %v completed!\n", hash)
			return status, nil
		case "FAILED":
			fmt.Printf("transaction: %v failed!\n", hash)
			return status, nil
		case "PENDING":
			fmt.Printf("transaction: %v pending!\n", hash)
			if i == retry-1 {
				return status, nil
			}
		case "DNE":
			return status, fmt.Errorf("transaction does not exist: %v", hash)
		default:
			return status, fmt.Errorf("unknown status: %v", status)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return "", fmt.Errorf("transaction monitoring failed after %v retry\n", retry)
}
