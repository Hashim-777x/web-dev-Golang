  // Phase4/Api.go

package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Api struct {
	addr string
	users []User
	mu sync.Mutex
}

 // slice struct to store the users in memory ... our DB 

// --------------------GET USERS HANDLER-----------------------------

func(a *Api) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

 w.Header().Set("Content-Type", "application/json") // Set the Content-Type header to application/json to indicate that the response will be in JSON format

 a.mu.Lock()
 usersCopy := make([]User, len(a.users))
 copy(usersCopy, a.users)
 a.mu.Unlock()

 err := json.NewEncoder(w).Encode(usersCopy) // Encode the Users slice as JSON and write it to the response writer
 if err != nil {            // returning the "Users" slice as JSON response to client
  http.Error(w, err.Error(), http.StatusInternalServerError) // If there is an error encoding the Users slice, return a 500 Internal Server Error response
  return
 }

// w.WriteHeader(http.StatusOK) // Set the response status code to 200 OK

}

// --------------------CREATE USER HANDLER-----------------------------

func (a *Api) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
 
  var payload User // Create a new instance of the User struct to hold the incoming request payload

  if err := json.NewDecoder(r.Body).Decode(&payload); err != nil { // Decode the incoming request body as JSON and store it in the payload variable
	http.Error(w, err.Error(), http.StatusBadRequest) // If there is an error decoding the request body, return a 400 Bad Request response
	return
  }

 U:= User{
	FirstName: payload.FirstName, // Create a new User instance with the decoded payload data
	LastName: payload.LastName,
  } 

  if err := a.Insertusers(U); err != nil { // Call the Insertusers function to insert the new user into the Users slice
	http.Error(w, err.Error(), http.StatusBadRequest) // If there is an error inserting the user, return a 400 Bad Request response
	return
  }
  w.WriteHeader(http.StatusCreated) // Set the response status code to 201 Created

}

func (a *Api) Insertusers(u User)error {
	if u.FirstName == "" || u.LastName == "" { // Check if the FirstName or LastName fields are empty
		return fmt.Errorf("first name and last name are required") // If either field is empty, return an error indicating that both fields are required
	}

 a.mu.Lock() // Lock the mutex to ensure that only one goroutine can access the Users slice at a time
 defer a.mu.Unlock() // Unlock the mutex after the user has been inserted

 for _, user := range a.users { // Iterate over the existing users in the Users slice

	if user.FirstName == u.FirstName && user.LastName == u.LastName { // Check if a user with the same FirstName and LastName already exists
		return fmt.Errorf("user already exists") // If a matching user is found, return an error indicating that the user already exists
	}
  }
 a.users = append(a.users, u) // Append the new user to the Users slice
 return nil // Return nil to indicate that the user was inserted successfully

}