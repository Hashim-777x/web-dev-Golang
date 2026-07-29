// making web server with different HTTP methods and URL paths 
// Using MUX for request routing and handling different HTTP methods and URL paths

package main 
 import(
	"net/http"
 )

 type Api struct {
	addr string
}

 func (a *Api) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

 func (a *Api) GetUserHandler(w http.ResponseWriter, r *http.Request) {

}

func main(){

 api := &Api{addr: ":8083"}
 
 
 server := &http.Server{ // create a new http.Server interface has two fields: Addr and Handler. 
	Addr: api.addr, //The Addr field specifies the address on which the server will listen for incoming requests, 
	Handler: mux, //and the Handler field specifies the handler that will handle incoming requests.
}

if err := server.ListenAndServe() ;err != nil {// start the server and listen for incoming requests
	panic(err) // if an error occurs while starting the server, panic and print the error message
}
fmt.Println("Server is running on port 8083")

mux := http.NewServeMux() // create a new ServeMux instance, which is a request router that matches incoming requests to their respective handlers based on the URL path.

 mux.HandleFunc("GET /users ", api.GetUserHandler) // register the GetUserHandler function to handle requests to the "GET/users" URL path
 mux.HandleFunc("POST /users", api.CreateUserHandler) // register the CreateUserHandler function to handle requests to the "POST/users" URL path

}