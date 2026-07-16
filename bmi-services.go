package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type BMIResponse struct {

	Weight float64 `json:"weight"`

	Height float64 `json:"height"`

	BMI float64 `json:"bmi"`

	Status string `json:"status"`

}



func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"BMI Calculator API is Running")

}



func calculateBMI(w http.ResponseWriter, r *http.Request){


	weightStr := r.URL.Query().Get("weight")

	heightStr := r.URL.Query().Get("height")



	weight, err := strconv.ParseFloat(weightStr,64)

	if err != nil {

		http.Error(w,"Invalid weight",400)

		return

	}



	height, err := strconv.ParseFloat(heightStr,64)

	if err != nil {

		http.Error(w,"Invalid height",400)

		return

	}



	if height <= 0 {

		http.Error(w,"Height must be greater than zero",400)

		return

	}



	bmi := weight / (height * height)



	status := ""


	switch {


	case bmi < 18.5:

		status = "Underweight"


	case bmi >= 18.5 && bmi < 25:

		status = "Normal Weight"


	case bmi >= 25 && bmi < 30:

		status = "Overweight"


	default:

		status = "Obese"

	}



	response := BMIResponse{

		Weight: weight,

		Height: height,

		BMI: bmi,

		Status: status,

	}



	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(response)


}



func main(){


	http.HandleFunc("/",home)


	http.HandleFunc("/bmi",calculateBMI)



	fmt.Println("BMI API running on port 8080")


	http.ListenAndServe(":8080",nil)

}