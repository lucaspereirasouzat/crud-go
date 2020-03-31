package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	// resposta para o front
	fmt.Fprintf(w, "bem vindo")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
