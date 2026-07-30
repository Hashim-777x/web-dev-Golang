  // Phase4/main.go

package main

import (
	"fmt"
	"net/http"
)

func main (){

  Api := &Api{addr: ":8084"} // Create an instance of the Api struct with the desired address
  mux := http.NewServeMux() // Create a new ServeMux to handle routing of incoming requests

  server := &http.Server{ // Create a new HTTP server with the specified address and handler
	Addr: Api.addr,
	Handler: mux,
  }

  mux.HandleFunc("GET /users", Api.GetUsersHandler) // Register the GetUsersHandler function to handle requests to the "/users" path
  mux.HandleFunc("POST /users", Api.CreateUserHandler) // Register the CreateUserHandler function to handle requests to the "/users" path
  
  fmt.Println("Server is running on port 8084") 
 if err := server.ListenAndServe(); err != nil { // start the server and listen for incoming requests, if there is an error, panic
		panic(err)           // last thing
	}
}
