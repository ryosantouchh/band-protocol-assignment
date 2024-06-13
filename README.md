### Prerequiste
- Go version 1.22  

For each task,I already implemented the logic and finished all of them. below is the intruction for running the code.  

### Task 1 - boss baby's revenge

```
cd bossbaby
go run main.go
```

```
input -> output
SRSSRRR -> "Good Boy"
RSSRR -> "Bad Boy"
SSSRRRRS -> "Bad Boy"
SRRSSR -> "Bad Boy"
SSRSRRR -> "Good Boy"
```
---

### Task 2 - superman's chicken rescue

```
cd superman
go run main.go
```

```
input -> output
[2,5,10,12,15] , 5 -> 2
[1,11,30,34,35,37] , 10 -> 4
```

---

### Task 3 - transaction broadcasting and monitoring client

```
cd tx-monitoring
go run main.go
```

After run the main file , we currently can create and get the transaction status.  

#### Create Transaction via cURL

```
 curl -X POST -H "Content-Type: application/json" -d '{"symbol": "ETH", "price": 4500, "timestamp": 1678912345}' http://localhost:8080/create-tx
```

##### Response  
```
{
  "tx_hash":"b7cbb7d4e084c467c7a303ebe50169d456a9488c2a0542f781c52a112aef3588"
}
```
  
Server Log Example for creating transaction:  
```
Created transaction : ba07e21b9101f4de8d64989c157be101f8d6c647921d9d6032610cec6554bcc4
Symbol : ETH
Price : 4500
Timestamp : 1678912345
```

#### Monitor Transaction via cURL  

- You need to use transaction from the response of create transaction as a path parameter

```
curl -X GET -H "Content-Type: application/json" http://localhost:8080/get-tx-status/b7cbb7d4e084c467c7a303ebe50169d456a9488c2a0542f781c52a112aef3588
```

##### Response

```
{
  "tx_status":"completed"
}
```
  
Server Log Example for PENDING status:  
```
start monitoring transaction: 73dc7eb7c6966bc122320d572ceee2d2cf6d98181f487c4982b9a1e69edb4019
transaction: 73dc7eb7c6966bc122320d572ceee2d2cf6d98181f487c4982b9a1e69edb4019 pending!
transaction: 73dc7eb7c6966bc122320d572ceee2d2cf6d98181f487c4982b9a1e69edb4019 pending!
```

Server Log Example for COMPLETED status:  
```
start monitoring transaction: 73dc7eb7c6966bc122320d572ceee2d2cf6d98181f487c4982b9a1e69edb4019
transaction: 73dc7eb7c6966bc122320d572ceee2d2cf6d98181f487c4982b9a1e69edb4019 completed!
```


#### We already got the response from server via the cURL above, however we can use it as a CLIENT to monitoring when someone try to create transaction or set transaction status via this application


