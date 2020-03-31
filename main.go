package main

// imports principais
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//http.HandleFunc("/", test)
	// iniciar servidor na porta localhost:1337
	println("Servidor iniciado na porta 1337")
	// http.ListenAndServe(":1337", nil)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", test)
	router.HandleFunc("/todo", TodoIndex)
	router.HandleFunc("/todo:{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":1337", router))
}

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
