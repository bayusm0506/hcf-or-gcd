package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bayusm0506/hcf-or-gcd/app/models"
)

// Func HCF
func Hcf(w http.ResponseWriter, r *http.Request) {
	hcf := models.Hcf{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&hcf); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var result int

	for i := 1; i <= hcf.Number1 && i <= hcf.Number2; i++ {
		// Check if is factor of both integers
		if hcf.Number1%i == 0 && hcf.Number2%i == 0 {
			result = i
		}
	}

	respondJSON(w, http.StatusCreated, result)
}
