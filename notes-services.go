package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)


type Note struct {

	ID int `json:"id"`

	Title string `json:"title"`

	Content string `json:"content"`

}


var notes []Note



func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"Notes API is Running")

}



// GET all notes
func getNotes(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(notes)

}



// CREATE note
func createNote(w http.ResponseWriter, r *http.Request){


	var note Note


	err := json.NewDecoder(r.Body).Decode(&note)


	if err != nil {

		http.Error(w,"Invalid JSON",400)

		return

	}



	note.ID = len(notes)+1


	notes = append(notes,note)



	w.Header().Set("Content-Type","application/json")


	json.NewEncoder(w).Encode(note)

}



// GET note by ID
func getNoteByID(w http.ResponseWriter, r *http.Request){


	idStr := strings.TrimPrefix(r.URL.Path,"/notes/")


	id,err := strconv.Atoi(idStr)


	if err != nil {

		http.Error(w,"Invalid ID",400)

		return

	}



	for _, note := range notes {


		if note.ID == id {


			json.NewEncoder(w).Encode(note)

			return

		}

	}



	http.Error(w,"Note not found",404)

}



// UPDATE note
func updateNote(w http.ResponseWriter, r *http.Request){


	idStr := strings.TrimPrefix(r.URL.Path,"/notes/")


	id,err := strconv.Atoi(idStr)


	if err != nil {

		http.Error(w,"Invalid ID",400)

		return

	}



	for i := range notes {


		if notes[i].ID == id {


			var updated Note


			json.NewDecoder(r.Body).Decode(&updated)



			notes[i].Title = updated.Title

			notes[i].Content = updated.Content



			json.NewEncoder(w).Encode(notes[i])


			return

		}

	}



	http.Error(w,"Note not found",404)

}



// DELETE note
func deleteNote(w http.ResponseWriter, r *http.Request){


	idStr := strings.TrimPrefix(r.URL.Path,"/notes/")


	id,err := strconv.Atoi(idStr)


	if err != nil {

		http.Error(w,"Invalid ID",400)

		return

	}



	for i := range notes {


		if notes[i].ID == id {


			notes = append(notes[:i],notes[i+1:]...)


			fmt.Fprintln(w,"Note deleted")


			return

		}

	}



	http.Error(w,"Note not found",404)

}



// Router
func notesHandler(w http.ResponseWriter,r *http.Request){


	switch r.Method {


	case http.MethodGet:


		if r.URL.Path == "/notes" {


			getNotes(w,r)


		}else{


			getNoteByID(w,r)

		}



	case http.MethodPost:


		createNote(w,r)



	case http.MethodPut:


		updateNote(w,r)



	case http.MethodDelete:


		deleteNote(w,r)



	default:


		http.Error(w,"Method not allowed",405)

	}

}



func main(){


	http.HandleFunc("/",home)


	http.HandleFunc("/notes",notesHandler)


	http.HandleFunc("/notes/",notesHandler)



	fmt.Println("Notes API running on port 8080")


	http.ListenAndServe(":8080",nil)

}