package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AndreiAlbert/brainf/generators"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{
			http.MethodPost,
		},
	})
	http.ListenAndServe(":8080", corsOpts.Handler(router))
}
