package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type EvenOddResponse struct {

	Number int `json:"number"`

	Result string `json:"result"`

}



func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"Even Odd Checker API is Running")

}



func checkEvenOdd(w http.ResponseWriter, r *http.Request){


	numberStr := r.URL.Query().Get("number")


	number, err := strconv.Atoi(numberStr)


	if err != nil {

		http.Error(w,"Invalid number",400)

		return

	}



	result := ""


	if number%2 == 0 {

		result = "Even"

	} else {

		result = "Odd"

	}



	response := EvenOddResponse{

		Number: number,

		Result: result,

	}



	w.Header().Set("Content-Type","application/json")


	json.NewEncoder(w).Encode(response)


}



func main(){


	http.HandleFunc("/",home)


	http.HandleFunc("/check",checkEvenOdd)


	fmt.Println("Even Odd API running on port 8080")


	http.ListenAndServe(":8080",nil)


}