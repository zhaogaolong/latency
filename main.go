package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/zhaogaolong/latency/pkg/network"
)

const (
	timeMsSuffix  = "ms"
	timeSecSuffix = "s"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func validationLatencyTime(t string) error {
	if strings.HasSuffix(t, timeMsSuffix) || strings.HasSuffix(t, timeSecSuffix) {
		return nil
	}
	return fmt.Errorf("%s invalide latency time, must like 10s or 10ms", t)
}

func Latency(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	latencyTime := ps.ByName("time")
	err := validationLatencyTime(latencyTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// to do set tc rule
	err = network.SetLatency(latencyTime, "eth0")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "set latency %s success\n", latencyTime)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/latency/:time", Latency)

	log.Println("start server :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
