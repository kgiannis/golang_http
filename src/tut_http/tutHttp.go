package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var workingDir string

type Person struct {
	Name string
	Surname string
	Age int
}

type Page struct {
	Message string
	Persons []Person
}

/* Get the working directory */
func getWorkingDir() string  {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return workingDir
}

func main() {
	/* Register HTTP handlers */
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/persons", personsHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	/* Get the template file */
	personsTemplate := template.Must(template.ParseFiles(getWorkingDir() + "/html/persons.html"))

	message := "Index Page"
	data := Page {
		Message: message,
	}

	personsTemplate.Execute(w, data)
}

func personsHandler(w http.ResponseWriter, r *http.Request) {
	/* Get the template file */
	personsTemplate := template.Must(template.ParseFiles(getWorkingDir() + "/html/persons.html"))

	message := "List of Persons"
	data := Page {
		Message: message,
		Persons: []Person {
			{Name: "John", Surname: "Doe", Age: 39},
			{Name: "Jane", Surname: "Doe", Age: 31},
		},
	}
	personsTemplate.Execute(w, data)
}