package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}

var students []Student

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Student API is Running")
}


// GET students
func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(students)

}


// POST student
func addStudent(w http.ResponseWriter, r *http.Request) {

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return

	}


	student.ID = len(students) + 1


	students = append(students, student)


	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(student)

}


func studentHandler(w http.ResponseWriter, r *http.Request) {


	switch r.Method {


	case http.MethodGet:

		getStudents(w,r)


	case http.MethodPost:

		addStudent(w,r)


	default:

		http.Error(w,"Method not allowed",405)

	}

}



func main(){

	http.HandleFunc("/",home)

	http.HandleFunc("/students",studentHandler)


	fmt.Println("Student API running on port 8080")


	http.ListenAndServe(":8080",nil)

}