package main

import (
	"net/http"
	"html/template"
	"fmt"
)
type Page struct{
		Name string
	}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	data := Page{
		Name: "John",
	}
	
	//fmt.Fprintln(w, "Welcome To Ascii Art Web")
	templ.Execute(w, data)
}

// func AboutHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w, "About")
// }
// func ContantHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w, "Contact Loaded Successfully.")
// }
func main() {

	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/About", AboutHandler)
	// http.HandleFunc("/Contant", ContantHandler)
	fmt.Println("server loading successfully!")
	http.ListenAndServe(":8080", nil)
}