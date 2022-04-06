package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"beererks.github.io/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//interate over all the mocks book
	for index, book := range mocks.Books {
		if book.Id == id {
			//delete and send book when id match dynamic
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
