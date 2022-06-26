package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleName).Methods("GET")
	router.HandleFunc("/bad", handleServerError).Methods("GET")
	router.HandleFunc("/data", handleData).Methods("POST")
	router.HandleFunc("/headers", handleHeader).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func handleName(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)["PARAM"]
	s := fmt.Sprintf("Hello, %s!", p)
	fmt.Fprint(w, s)
}

func handleServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}
	fmt.Println(string(data))
	param := fmt.Sprintf("I got message:\n%s", string(data))

	w.Write([]byte(param))
}

func handleHeader(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}
	result := strconv.Itoa(a + b)
	w.Header().Set("a+b", result)
}
