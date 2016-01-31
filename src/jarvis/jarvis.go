package main

import (
	"net/http"
	"os"
	"jarvis/handlers"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(handlers.JustFilesFilesystem{http.Dir(os.Getenv("GOPATH") + "/public")})
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", Index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}