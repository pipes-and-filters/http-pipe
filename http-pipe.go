package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/pipes-and-filters/filters"

	"log"
)

var (
	chainFile string
	port      string
	chain     filters.Chain
)

func init() {
	flag.StringVar(&pxec, "chain-file", os.Getenv("HTTP_PIPE_CHAIN"), "Chain file for http pipe")
	flag.StringVar(&port, "port", os.Getenv("PORT"), "Port to serve over http")
}

func main() {
	flag.Parse()
	chain, err := filters.ChainFile("chain.yml")
	if err != nil {
		log.Fatal(err)
	}
	_, err := chain.Exec()
	if err != nil {
		log.Fatal(err)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), http.HandlerFunc(FilterHandler))
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	e, err := chain.Exec()
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Error(err)
	}
	e.SetInput(r.Body)
	e.SetOutput(w)
	err = e.Run()
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Error(err)
	}
	go logErrors(e.Errors())
}

func logErrors(es []error) {
	for _, e := range es {
		log.Error(e)
	}
}
