package main

import (
	"net/http"

	"github.com/MurilloSSJ/goflow/api/modules/dags"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/dags", dags.DagMux())
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}