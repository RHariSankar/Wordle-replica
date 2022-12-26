package routes

import (
	"backend/words"
	"encoding/json"
	"net/http"
)

type IsValidResponse struct {
	Word    string `json:"word"`
	IsValid bool   `json:"isValid"`
}

func IsValidWord(w http.ResponseWriter, r *http.Request) {
	var trie = words.GetInstance()
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	word, ok := requestBody["word"]
	if !ok {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	isValid := trie.Find(word.(string))
	response := IsValidResponse{
		Word:    word.(string),
		IsValid: isValid,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "couldn't marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "couldn't write response", http.StatusInternalServerError)
		return
	}
}
