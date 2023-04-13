package proxy

import (
	"fmt"
	"sync"
	"testing"
)

func TestProxy(t *testing.T) {
	nginxServer := NewNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	wg := sync.WaitGroup{}

	var httpCode int
	var err error

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			httpCode, err = nginxServer.handleRequest(appStatusURL, "GET")
			fmt.Printf("\nUrl: %s\nHttpCode: %d\nErr: %v\n", appStatusURL, httpCode, err)
		}()
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			httpCode, err = nginxServer.handleRequest(createuserURL, "POST")
			fmt.Printf("\nUrl: %s\nHttpCode: %d\nErr: %v\n", createuserURL, httpCode, err)
		}()
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			httpCode, err = nginxServer.handleRequest(createuserURL, "GET")
			fmt.Printf("\nUrl: %s\nHttpCode: %d\nErr: %v\n", createuserURL, httpCode, err)
		}()
	}

	wg.Wait()
}
