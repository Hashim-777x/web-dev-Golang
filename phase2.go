  //phase2.go
  
// A simple web server that listens on port 8080
// which response for specific path

// a web server and a handler **with-OUT using MUX** for request routing.

package main 

import (
	"net/http"
)

type Server struct {
	addr string // stores the address on which the server will listen for incoming requests
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) { // type Struct which implements the ServeHTTP method, which is required by the http.Handler interface. This method is called for each incoming HTTP request.

	switch r.Method { // check the what HTTP method of the incoming request is (GET, POST, etc.) and handle it accordingly

//-------------GET method handling------------------

		case http.MethodGet : // ** if the method is GET **

		    switch r.URL.Path { // another switch to check the URL path of the incoming request
		     case "/":
				w.Write([]byte("Get Method response from the server!")) // respond with "Hello, World!" for the root path
		     case "/index":
				w.Write([]byte("getting the index page from the server!")) // respond with "Hello, World!" for the root path}
		    }

//-------------POST method handling-----------------		

		case http.MethodPost : // ** if the method is POST **

		   w.Write([]byte("Post Method response from the server!.. posted successfully")) // respond with message indicating that the POST request was successful
		   return http.StatusCreated // respond with 201 Created status code to indicate that the resource was successfully created

	}	
}


func main() {

 S := &Server{addr: ":8082"}

 if err := http.ListenAndServe(s.addr, s); err != nil { //takes address and who handles the request
		panic(err)
	}
fmt.Println("Server is running on port 8082")

// handling diffrent HTTP methods without using MUx for request routing.
//just using switch case to handle different HTTP methods and URL paths and respond accordingly.

}