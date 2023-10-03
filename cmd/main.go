package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	data     string
	dataLock sync.RWMutex
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	router.HandleFunc("/get", GetHandler).Methods("GET")

	router.HandleFunc("/post", PostHandler).Methods("POST")

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	dataLock.RLock()
	defer dataLock.RUnlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data: %s", data)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	dataLock.Lock()
	defer dataLock.Unlock()

	data = string(body)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data received and stored")
}
