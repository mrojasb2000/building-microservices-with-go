package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request 1", http.StatusBadRequest)
		return
	}
	log.Printf("data:%v", body)

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	log.Printf("err:%v", err)
	if err != nil {
		http.Error(w, "Bad request 2", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	callback := r.URL.Query().Get("callback")
	if callback != "" {
		r.Header.Set("Content-Type", "application/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, string(data))
	} else {
		fmt.Fprintf(w, string(data))
	}
}
