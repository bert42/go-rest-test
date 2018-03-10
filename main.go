package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"net/http"
	"fmt"
	"strconv"
	"flag"
	"time"
)

const PORT = 3000;

func main() {
	var dir string;

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the ./static dir");
    flag.Parse();

	router := mux.NewRouter();

	router.HandleFunc("/ping", PingHandler);

	// serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))));
	router.HandleFunc("/", IndexHandler);

	fmt.Println("Starting http server on http://localhost:", PORT);

	srv := &http.Server {
        Handler:      handlers.LoggingHandler(os.Stdout, router),
        Addr:         ":"+strconv.Itoa(PORT),
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,

	};

	srv.ListenAndServe();
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r);

    w.WriteHeader(http.StatusOK);

    fmt.Fprintf(w, "Pong")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/index.html", http.StatusMovedPermanently);
}
