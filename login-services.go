package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type LoginRequest struct {

	Username string `json:"username"`

	Password string `json:"password"`

}


type LoginResponse struct {

	Message string `json:"message"`

	Status string `json:"status"`

}



func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"Login API is Running")

}



func login(w http.ResponseWriter, r *http.Request){


	if r.Method != http.MethodPost {

		http.Error(w,"Only POST method allowed",405)

		return

	}


	var user LoginRequest


	err := json.NewDecoder(r.Body).Decode(&user)


	if err != nil {

		http.Error(w,"Invalid JSON",400)

		return

	}



	response := LoginResponse{}



	// Demo authentication

	if user.Username == "admin" && user.Password == "1234" {


		response.Message = "Login Successful"

		response.Status = "success"


	}else{


		response.Message = "Invalid Username or Password"

		response.Status = "failed"

	}



	w.Header().Set("Content-Type","application/json")


	json.NewEncoder(w).Encode(response)


}




func main(){


	http.HandleFunc("/",home)


	http.HandleFunc("/login",login)


	fmt.Println("Login API running on port 8080")


	http.ListenAndServe(":8080",nil)


}