package proxy

import "errors"

type Application struct{}

func (a *Application) handleRequest(url, method string) (code int, err error) {
	if url == "/app/status" && method == "GET" {
		return StatusOK, nil
	}

	if url == "/create/user" && method == "POST" {
		return StatusCreated, nil
	}

	return StatusNotFound, errors.New("Not Found!")
}
