package main

import(
	"fmt"
	"net/http"
	"log"
)

func main(){
	fileServer := http.FileServer(http.Dir("/static"))

	http.Handle("/",fileServer)
	http.HandlerFunc("/form",formHandler)
	http.HandlerFunc("/home",homeHandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
}