package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

const PORT = "3000";

func main() {
	router := mux.NewRouter();

	router.HandleFunc("/ping", PingHandler);

	log.Fatal(http.ListenAndServe(":"+PORT, router));
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r);

    w.WriteHeader(http.StatusOK);

    fmt.Fprintf(w, "Pong")
}
