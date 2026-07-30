   // phase1.go

// A very simple web server that listens on port 8080
// and responds with "Hello, World!" to any incoming HTTP request.

package main

import (
	"net/http"
)

type Server struct {
	addr string // stores the address on which the server will listen for incoming requests
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
  // ResponseWriter write back to the client, Request contains all the information about the incoming request.


func main() {
	s := &Server{addr: ":8081"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port 8081")
}
