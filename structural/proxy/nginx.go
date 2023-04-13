package proxy

import (
	"errors"
	"sync"
)

const (
	StatusOK        = 200
	StatusCreated   = 201
	StatusForbidden = 403
	StatusNotFound  = 404
)

type Server interface {
	handleRequest(url, method string) (code int, err error)
}

type NginxServer struct {
	app    *Application
	maxReq int
	rates  map[string]int
	mu     sync.Mutex
}

func NewNginxServer() NginxServer {
	return NginxServer{
		app:    &Application{},
		maxReq: 2,
		rates:  make(map[string]int),
		mu:     sync.Mutex{},
	}
}

func (n *NginxServer) handleRequest(url, method string) (code int, err error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.checkRates(url) {
		return StatusForbidden, errors.New("Forbidden!")
	}

	return n.app.handleRequest(url, method)
}

func (n *NginxServer) checkRates(url string) bool {
	if n.rates[url] >= n.maxReq {
		return false
	}

	if n.rates[url] == 0 {
		n.rates[url] = 1
	} else {
		n.rates[url] += 1
	}

	return true
}
