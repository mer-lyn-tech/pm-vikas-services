package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type TemperatureResponse struct {

	From string `json:"from"`

	To string `json:"to"`

	Input float64 `json:"input"`

	Result float64 `json:"result"`

}



func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"Temperature Converter API is Running")

}



func convertTemperature(w http.ResponseWriter, r *http.Request){


	valueStr := r.URL.Query().Get("value")

	from := r.URL.Query().Get("from")

	to := r.URL.Query().Get("to")



	value, err := strconv.ParseFloat(valueStr,64)


	if err != nil {

		http.Error(w,"Invalid temperature value",400)

		return

	}



	var result float64



	switch {


	// Celsius to Fahrenheit
	case from=="celsius" && to=="fahrenheit":

		result = (value * 9 / 5) + 32



	// Fahrenheit to Celsius
	case from=="fahrenheit" && to=="celsius":

		result = (value - 32) * 5 / 9



	// Same unit
	case from==to:

		result = value



	default:

		http.Error(w,"Invalid conversion type",400)

		return

	}



	response := TemperatureResponse{

		From: from,

		To: to,

		Input: value,

		Result: result,

	}



	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(response)


}



func main(){


	http.HandleFunc("/",home)


	http.HandleFunc("/convert",convertTemperature)



	fmt.Println("Temperature API running on port 8080")


	http.ListenAndServe(":8080",nil)


}