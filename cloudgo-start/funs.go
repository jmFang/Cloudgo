package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "welcome!\n")
}

var bookstore = make(map[string]*Book)

func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := []*Book{}
	for _, book := range bookstore {
		books = append(books, book)
	}
	response := &JsonResponse{Data: &books}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isbn")
	book, ok := bookstore[isdn]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: "Record Not Found"}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}

	response := JsonResponse{Data: book}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
