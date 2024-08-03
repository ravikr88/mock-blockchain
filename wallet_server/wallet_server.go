package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

const tempDir = "templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port: port, gateway: gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, err := template.ParseFiles(path.Join(tempDir, "index.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		log.Printf("Error: Invalid HTTP Method")
		http.Error(w, "Invalid HTTP Method", http.StatusMethodNotAllowed)
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	addr := "0.0.0.0:" + strconv.Itoa(int(ws.Port()))
	log.Printf("Listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
