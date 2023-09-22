package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AndreiAlbert/brainf/generators"
	"github.com/gorilla/mux"
)

type request struct {
	Code string `json:"code"`
}

type response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Fatal(err.Error())
	}
	g := generators.New(req.Code)
	result, err := g.Evaluate()
	if err != nil {
		response := response{Error: err.Error()}
		jsonResponse, jsonError := json.Marshal(response)
		if jsonError != nil {
			fmt.Fprint(w, jsonError)
			return
		}
		w.WriteHeader(400)
		w.Write(jsonResponse)
		return
	}
	response := response{Result: result.String()}
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		fmt.Fprint(w, jsonError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/processRequest", processRequest).Methods("POST")
	http.ListenAndServe(":8080", router)
}
