package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
)

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func main() {
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}

	threadNumber := 20
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
				res, err := http.Get("http://" + ip + ":30001/demo")
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
