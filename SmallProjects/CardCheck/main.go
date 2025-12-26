package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"small_go_projects/CardCheck/luhn"
)

type Card struct {
	Number string `json:"number"`
}

func creditCardValidator(w http.ResponseWriter, r *http.Request) {
	var card Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid := luhn.LuhnAlgorithm(card.Number)
	response := map[string]bool{"valid": isValid}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	args := os.Args
	port := args[1]

	http.HandleFunc("/", creditCardValidator)
	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
