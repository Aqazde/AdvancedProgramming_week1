package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func fetchData() map[string]interface{} {
	data := map[string]interface{}{
		"user":     "Adiqde",
		"position": "owner",
		"ID":       1,
	}
	return data
}

func handleGetData(w http.ResponseWriter, r *http.Request) {
	data := fetchData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/getdata", handleGetData)
	http.HandleFunc("/", serveHTML)

	fmt.Println("Server listening on port 1997")
	log.Fatal(http.ListenAndServe(":1997", nil))
}
