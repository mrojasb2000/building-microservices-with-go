package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	//"io/ioutil"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	ID      int    `json:"id,string"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func helloWorldRequestHandler(w http.ResponseWriter, r *http.Request) {
	/*
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}*/

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	/*
		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}*/

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{
		Message: "Hello World!",
		Author:  "Mavro R0j4s",
		Date:    time.Now().String(),
		ID:      99,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

	/*
		data, err := json.Marshal(response)
		if err != nil {
			panic("Ooops")
		}
		fmt.Fprint(w, string(data))*/
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	http.HandleFunc("/", helloWorldRequestHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
