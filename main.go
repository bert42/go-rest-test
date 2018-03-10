package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"net/http"
	"fmt"
)

const PORT = "3000";

func main() {
	router := mux.NewRouter();

	router.HandleFunc("/ping", PingHandler);

	fmt.Print("Starting http server on http://localhost:"+PORT+"\n");
	http.ListenAndServe(":"+PORT, handlers.LoggingHandler(os.Stdout, router));
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r);

    w.WriteHeader(http.StatusOK);

    fmt.Fprintf(w, "Pong")
}
