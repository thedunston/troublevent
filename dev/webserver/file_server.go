package main

/* Tutorial from: https://zetcode.com/golang/http-server/ */

import (

	"fmt"
	"log"
	"io"
	"net/http"

)

func main() {

	// Create a web server listening on Port 8082
	// The HandleFunc registers the handler function for the given URL pattern.
	// web root
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello",  func(w http.ResponseWriter, r *http.Request) {

		io.WriteString(w, "Hello there\n")
	})

	/**
	// Get the name parameter
	keys, ok := r.URL.Query()["name"]

	// Default username printed to the screen.
	name := "guest"

	if ok {

		// Query returned if user inputs a name
		// http://localhost:8082/?name=duane
		name = keys[0]

	}

	fmt.Fprintf(w, "hello %s\n", name)
	*/
	fmt.Println("Server started at port 8082")
	// Opens the port and listens for incoming connections.
	log.Fatal(http.ListenAndServe(":8082", nil))

}

/**func HelloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, there\n")

}
*/

/*func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[2])

}*/
