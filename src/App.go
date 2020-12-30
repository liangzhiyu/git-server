package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r:= mux.NewRouter()
	r.HandleFunc("/{user}/{repo}/info/refs", handleInfoRefs).Queries().Methods("GET")
	r.HandleFunc("/{user}/{repo}/git-upload-pack", handleUploadPack).Methods("POST")
	r.HandleFunc("/{user}/{repo}/git-receive-pack", handleReceivePack).Methods("POST")
	r.Handle("/debug/pprof/", http.DefaultServeMux)

	http.ListenAndServe(":8989", r)
}

func handleReceivePack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleReceivePack")
}

func handleUploadPack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleUploadPack")
}

func handleInfoRefs(w http.ResponseWriter, r *http.Request) {
	service := r.URL.Query().Get("service")
	fmt.Println("handleInfoRefs", service)
}
