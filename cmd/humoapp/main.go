package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	ip := "localhost"
	port := ":8000"
	adress := ip + port

	router := mux.NewRouter()
	router.HandleFunc("/calculator/sum", Sum).Methods(http.MethodPost)
	router.HandleFunc("/calculator/substraction", Substraction).Methods(http.MethodPost)
	router.HandleFunc("/calculator/multiplication", Multiplication).Methods(http.MethodPost)
	router.HandleFunc("/calculator/division", Division).Methods(http.MethodPost)

	err := http.ListenAndServe(adress, router)
	if err != nil {
		log.Println("listen and serve error", err)
		return
	}
}

type CalculatorRequest struct {
	NumberFirst  int `json:"number_first"`
	NumberSecond int `json:"number_second"`
}

func Sum(response http.ResponseWriter, request *http.Request) {
	var numbers CalculatorRequest

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bytes, &numbers)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result := numbers.NumberFirst + numbers.NumberSecond
	resp := strconv.Itoa(result)
	response.Write([]byte(resp))
}

func Substraction(response http.ResponseWriter, request *http.Request) {
	var numbers CalculatorRequest

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bytes, &numbers)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result := numbers.NumberFirst - numbers.NumberSecond
	resp := strconv.Itoa(result)
	response.Write([]byte(resp))
}

func Multiplication(response http.ResponseWriter, request *http.Request) {
	var numbers CalculatorRequest

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bytes, &numbers)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result := numbers.NumberFirst * numbers.NumberSecond
	resp := strconv.Itoa(result)
	response.Write([]byte(resp))
}

func Division(response http.ResponseWriter, request *http.Request) {
	var numbers CalculatorRequest

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bytes, &numbers)
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result := numbers.NumberFirst / numbers.NumberSecond
	resp := strconv.Itoa(result)
	response.Write([]byte(resp))
}
