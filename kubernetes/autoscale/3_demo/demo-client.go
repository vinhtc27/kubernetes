package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	loadbalancerExternalIP := "137.184.250.101"
	threadNumber := 100
	var channel = make(chan int, threadNumber+1)
	var waitGroup sync.WaitGroup

	waitGroup.Add(threadNumber)
	for thread := 0; thread < threadNumber; thread++ {
		go func() {
			for {
				call, stillCall := <-channel
				if !stillCall {
					waitGroup.Done()
					return
				}
				res, err := http.Get("http://" + loadbalancerExternalIP + "/demo")
				if err != nil {
					fmt.Println(err)
				}

				body, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
				}

				var decodedBody = make(map[string]any)
				err = json.Unmarshal(body, &decodedBody)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(call, "=> ", decodedBody["data"])
			}
		}()
	}
	for count := 1; count <= 1000000; count++ {
		channel <- count
	}
	close(channel)
	waitGroup.Wait()
}
