package controllers

import (
	"fmt"
	"net/http"
)

// Home used for / route
func (s *Server) Home(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Welcome to the api")
}
