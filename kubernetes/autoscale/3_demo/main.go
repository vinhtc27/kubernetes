package main

import (
	"fmt"
	"net/http"
	"time"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	var count uint64 = 0
	startTime := time.Now()
	for {
		count++
		if count >= 10000000000 {
			break
		}
	}
	fmt.Fprintf(w, "%s\n", time.Since(startTime))
}
func main() {
	http.HandleFunc("/", Handle)
	http.ListenAndServe("localhost:3000", nil)
}
