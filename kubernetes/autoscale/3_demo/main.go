package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func ResponseWrite(w http.ResponseWriter, responseCode int, responseData interface{}) {
	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	// Write JSON to Response
	json.NewEncoder(w).Encode(responseData)
}

func ResponseWithData(w http.ResponseWriter, data interface{}) {
	var response struct {
		Status  bool        `json:"status"`
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	// Set Response Data
	response.Status = true
	response.Code = "200"
	response.Message = "Request successfully"
	response.Data = data

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusCreated, response)
}

func Demo(w http.ResponseWriter, req *http.Request) {
	threadNumber := 100
	var channel = make(chan int, threadNumber+1)
	var waitGroup sync.WaitGroup

	waitGroup.Add(threadNumber)
	var startCrawlTime = time.Now()
	for thread := 0; thread < threadNumber; thread++ {
		go func() {
			for {
				_, stillCount := <-channel
				if !stillCount {
					waitGroup.Done()
					return
				}
				i := 0
				for {
					i++
					if i >= 300 {
						break
					}
				}
			}
		}()
	}
	rand.Seed(time.Now().UnixNano())
	for count := 1; count <= rand.Intn(140000-100000)+100000; count++ {
		channel <- count
	}

	close(channel)
	waitGroup.Wait()
	timeDiff := time.Since(startCrawlTime)
	ResponseWithData(w, fmt.Sprintf("Response time: %s", timeDiff))

}
func main() {
	http.HandleFunc("/demo", Demo)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
