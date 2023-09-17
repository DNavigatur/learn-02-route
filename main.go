package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmp template.Template

func init() {
	tmp = *template.Must(template.ParseFiles("something.gohtml"))
}

type book struct {
	Name          string
	Author        string
	NumberOfpages int
}

func main() {

	// Define handler functions for different routes
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		data1 := book{
			Name:          "Index Book",
			Author:        "Micheal Jackson",
			NumberOfpages: 8,
		}
		err := tmp.Execute(w, data1)
		if err != nil {
			log.Fatalln(err)
		}
	}

	dogHandler := func(w http.ResponseWriter, r *http.Request) {
		data2 := book{
			Name:          "Dog Book",
			Author:        "Doja Cat",
			NumberOfpages: 5,
		}
		err := tmp.Execute(w, data2)
		if err != nil {
			log.Fatalln(err)
		}
	}

	meHandler := func(w http.ResponseWriter, r *http.Request) {
		data3 := book{
			Name:          "Henry Book",
			Author:        "Eze Henry",
			NumberOfpages: 3,
		}
		err := tmp.Execute(w, data3)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Register the handlers for specific routes using http.HandleFunc
	http.HandleFunc("/", indexHandler)   //index route
	http.HandleFunc("/dog/", dogHandler) //dog route
	http.HandleFunc("/me/", meHandler)   //me route

	//ListenAndServe on port ":8080" using the default ServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
