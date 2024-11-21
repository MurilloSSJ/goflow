package dags

import "net/http"

func DagMux() http.Handler {
	useMux := http.NewServeMux()
	useMux.Handle("GET /" , http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
	return useMux

}