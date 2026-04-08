package main

import(
	"fmt"
	"net/http"
	"log"
)

// fileServer for getting all files from static folder
func main(){
	fileServer := http.FileServer(http.Dir("/static"))


	// Serve static files from the root path (HTML, CSS, JS, images)
	http.Handle("/", fileServer)

	// Handle form submissions at /form
	http.HandleFunc("/form", formHandler)

	// Serve the home page at /home
	http.HandleFunc("/home", homeHandler)

	//Start http server on port 8080
	fmt.Printf("starting server on port 8080\n")
	if err := http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
}