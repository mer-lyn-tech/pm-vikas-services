package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type SalaryResponse struct {

	BasicSalary float64 `json:"basic_salary"`

	HRA float64 `json:"hra"`

	DA float64 `json:"da"`

	TotalSalary float64 `json:"total_salary"`

}


func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"Employee Salary Calculator API is Running")

}



func calculateSalary(w http.ResponseWriter, r *http.Request){


	salaryStr := r.URL.Query().Get("salary")


	salary, err := strconv.ParseFloat(salaryStr,64)


	if err != nil {

		http.Error(w,"Invalid salary value",400)

		return

	}



	// Business Logic

	// HRA = 20% of basic salary

	// DA = 10% of basic salary


	hra := salary * 0.20

	da := salary * 0.10


	total := salary + hra + da



	response := SalaryResponse{

		BasicSalary: salary,

		HRA: hra,

		DA: da,

		TotalSalary: total,

	}



	w.Header().Set("Content-Type","application/json")


	json.NewEncoder(w).Encode(response)


}



func main(){


	http.HandleFunc("/",home)


	http.HandleFunc("/salary",calculateSalary)



	fmt.Println("Salary API running on port 8080")


	http.ListenAndServe(":8080",nil)


}