package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"beererks.github.io/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	//read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//interate over all the mocks book
	for _, book := range mocks.Books {
		if book.Id == id {
			//if ids book are equal send book as response
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
