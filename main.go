package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonStruct struct {
	Env     string `json:"env"`
	Version string `json:"version"`
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Println("server started")
	http.ListenAndServe(":8000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var buildjson jsonStruct
	buildjson.Env = "some value"
	buildjson.Version = "69.420"
	jsonresp, err := json.Marshal(buildjson)
	if err != nil {
		log.Println("problem jsoning")
	}
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonresp)
}
