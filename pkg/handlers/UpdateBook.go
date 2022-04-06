package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"beererks.github.io/pkg/mocks"
	"beererks.github.io/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)
	//interate over all the mocks book
	for index, book := range mocks.Books {
		if book.Id == id {
			//Update and send book when id match dynamic
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
