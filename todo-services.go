package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)


type Todo struct {

	ID int `json:"id"`

	Title string `json:"title"`

	Completed bool `json:"completed"`

}


var todos []Todo



func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"To-Do API is Running")

}



// GET all todos
func getTodos(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(todos)

}



// CREATE todo
func createTodo(w http.ResponseWriter, r *http.Request){


	var todo Todo


	err := json.NewDecoder(r.Body).Decode(&todo)


	if err != nil {

		http.Error(w,"Invalid JSON",400)

		return

	}



	todo.ID = len(todos)+1


	todos = append(todos,todo)


	w.Header().Set("Content-Type","application/json")


	json.NewEncoder(w).Encode(todo)

}



// UPDATE todo
func updateTodo(w http.ResponseWriter, r *http.Request){


	idStr := strings.TrimPrefix(r.URL.Path,"/todos/")


	id,err := strconv.Atoi(idStr)


	if err != nil {

		http.Error(w,"Invalid ID",400)

		return

	}



	for i := range todos {


		if todos[i].ID == id {


			var updated Todo


			json.NewDecoder(r.Body).Decode(&updated)



			todos[i].Title = updated.Title

			todos[i].Completed = updated.Completed



			json.NewEncoder(w).Encode(todos[i])

			return

		}

	}



	http.Error(w,"Todo not found",404)

}



// DELETE todo
func deleteTodo(w http.ResponseWriter, r *http.Request){


	idStr := strings.TrimPrefix(r.URL.Path,"/todos/")


	id,err := strconv.Atoi(idStr)


	if err != nil {

		http.Error(w,"Invalid ID",400)

		return

	}



	for i := range todos {


		if todos[i].ID == id {


			todos = append(todos[:i],todos[i+1:]...)


			fmt.Fprintln(w,"Todo deleted")

			return

		}

	}



	http.Error(w,"Todo not found",404)

}



// Router
func todoHandler(w http.ResponseWriter,r *http.Request){


	switch r.Method {


	case http.MethodGet:


		if r.URL.Path == "/todos" {

			getTodos(w,r)

		}else{

			http.Error(w,"Invalid route",404)

		}



	case http.MethodPost:


		createTodo(w,r)



	case http.MethodPut:


		updateTodo(w,r)



	case http.MethodDelete:


		deleteTodo(w,r)



	default:


		http.Error(w,"Method not allowed",405)

	}

}




func main(){


	http.HandleFunc("/",home)


	http.HandleFunc("/todos",todoHandler)


	http.HandleFunc("/todos/",todoHandler)



	fmt.Println("Todo API running on port 8080")


	http.ListenAndServe(":8080",nil)

}