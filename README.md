# detailed Readme comming soon !!! please wait )
# Some bug fixing is being done !!! 
## syncing of txns among nodes ... ToDo !!
- Built a local blockchain network in Golang with features including Proof of Work, digital signatures, public and private keys, mining, wallet generation, and transaction verification, transferring cryptocurrency, synchronizing transactions, and maintaining a distributed ledger. Incorporated features such as block creation, chain validation, consensus mechanisms, and transaction broadcasting and an UI for handy use.


  # How a user A access to blockchain node ?
  - User A connects to wallet server running on port 8080 which connects to node server on port 9000

User-A  < == req/res === > 0.0.0.0:8080(wallet server) < === req/res === > 0.0.0.0:9000(node server)






# How to Run the nodes and wallet server ?
- node running on port 9000 (default) is a miner !!
- any wallet transferring money calls an api to this miner node
- 


- run a blockchain node server at specified port in _different terminal_ windows
- each node is running on respective port locally

```bash
// from blockchain_server directory run 
go run main.go blockchain_server.go  // default port is 9000 which is a miner
go run main.go blockchain_server.go -port 9001 // another node 
go run main.go blockchain_server.go -port 9002 // another node
```

### access the running blockchain server on 0.0.0.0:9000 (default port) on a browser window

<img width="574" alt="Screenshot 2024-08-05 at 7 14 12 PM" src="https://github.com/user-attachments/assets/b073698b-aa4f-4ff1-b94a-f705b38481b0">

- node running on port 9001

<img width="574" alt="Screenshot 2024-08-05 at 7 14 16 PM" src="https://github.com/user-attachments/assets/83b3a216-5eb0-47d0-aa9d-5f95779d3a11">

### to run wallet_server 
```bash
// from wallet_server directory run
go run main.go wallet_server.go -gateway = "http://127.0.0.1:9000"
```
### access the wallet UI at http://0.0.0.0:8080/
<img width="574" alt="Screenshot 2024-08-05 at 7 21 08 PM" src="https://github.com/user-attachments/assets/3242578d-35ec-4c0b-9408-2cf2649578a9">

### run other wallets by chainging the port and connecting to different node runnig on port 9001

```bash
go run main.go wallet_server.go -port 8080 -gateway http://127.0.0.1.9001
```










