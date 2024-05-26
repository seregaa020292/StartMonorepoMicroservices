package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"seregaa020292/StartMonorepoMicroservices/fines/controlling/internal/hi"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write(hi.Message())
	})
	mux.HandleFunc("GET /health/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("GET /ping/{$}", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(fmt.Sprintf("%s/ping", os.Getenv("MONITORING_ADDR")))
		if err != nil {
			w.Write([]byte(errors.WithStack(err).Error()))
			return
		}

		defer resp.Body.Close()

		buf := bytes.Buffer{}
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		log.Println(buf.String())
		buf.WriteTo(w)
	})

	log.Println("Starting server controlling")
	http.ListenAndServe(net.JoinHostPort("", "8080"), mux)
}
