package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/MurilloSSJ/goflow/api/modules/dags"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/dags", dags.DagMux())
	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}