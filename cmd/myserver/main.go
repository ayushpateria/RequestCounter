package main

import (
	"net/http"
	"strconv"
	
	"github.com/ayushpateria/RequestCounter/pkg/counter"
)

var c *counter.Counter

func withCounter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.Inc()
		next.ServeHTTP(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

func reportCount(w http.ResponseWriter, r *http.Request) {
	resp := strconv.FormatUint(c.Value(), 10)
	w.Write([]byte(resp))
}

func reportQps(w http.ResponseWriter, r *http.Request) {
	resp := strconv.FormatUint(c.Qps(), 10)
	w.Write([]byte(resp))
}

func main() {
	c = counter.NewCounter()
	http.HandleFunc("/count", reportCount)
	http.HandleFunc("/qps", reportQps)
	http.HandleFunc("/", withCounter(home))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
