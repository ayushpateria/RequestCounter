package main

import (
  "net/http"
  "strconv"
)

var c *counter

func withCounter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.inc()
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

func reportStats(w http.ResponseWriter, r *http.Request) {
	resp := strconv.FormatUint(c.value(), 10)
	w.Write([]byte(resp))
}

func main() {
  c = newCounter()
  http.HandleFunc("/count", reportStats)
  http.HandleFunc("/", withCounter(home))
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}