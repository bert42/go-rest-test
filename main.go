package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"net/http"
	"fmt"
	"strconv"
)

const PORT = 3000;

func main() {
	router := mux.NewRouter();

	router.HandleFunc("/ping", PingHandler);

	fmt.Println("Starting http server on http://localhost:", PORT);
	http.ListenAndServe(":"+strconv.Itoa(PORT), handlers.LoggingHandler(os.Stdout, router));
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r);

    w.WriteHeader(http.StatusOK);

    fmt.Fprintf(w, "Pong")
}
