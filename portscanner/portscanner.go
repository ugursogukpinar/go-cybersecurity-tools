package portscanner

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

type PortNumber = uint16

var ErrInvalidRange = errors.New("invalid range, end must be greater than start")

func IsPortOpen(host string, port PortNumber) bool {
	conn_str := fmt.Sprintf("%s:%d", host, port)

	duration, _ := time.ParseDuration("2s")
	conn, err := net.DialTimeout("tcp", conn_str, duration)
	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}

func GetOpenPorts(host string, start, end PortNumber) ([]PortNumber, error) {
	fmt.Println(host, start, end)
	if end < start {
		return nil, ErrInvalidRange
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	openPorts := []PortNumber{}

	for i := start; i <= end; i++ {
		wg.Add(1)

		go func(p PortNumber) {
			defer wg.Done()
			isOpen := IsPortOpen(host, p)

			if isOpen {
				mutex.Lock()
				openPorts = append(openPorts, p)
				mutex.Unlock()
			}
		}(i)
	}

	wg.Wait()

	return openPorts, nil
}
