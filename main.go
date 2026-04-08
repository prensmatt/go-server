package main

import(
	"fmt"
	"net/http"
	"log"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm();err != nil{
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request successful")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	if firstName == "" || lastName == "" || subject == "" || message == "" {
        http.Error(w, "please fill in all required fields", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "FirstName: %s\nLasttName: %s\nEmail: %s\nPhone: %s\nSubject: %s\nMessage: %s", firstName,lastName,email,phone, subject, message)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/home"{
		http.Error(w, "404 not found",http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"method is not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"home")
}

// fileServer for getting all files from static folder
func main(){
	fileServer := http.FileServer(http.Dir("./static"))


	// Serve static files from the root path (HTML, CSS, JS, images)
	http.Handle("/", http.StripPrefix("/", fileServer))

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