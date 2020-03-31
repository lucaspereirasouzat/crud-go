package main

// imports principais
import (
	"github.com/gorilla/mux"
	""
)

func routes() {
	//http.HandleFunc("/", test)
	// iniciar servidor na porta localhost:1337
	println("Servidor iniciado na porta 1337")
	// http.ListenAndServe(":1337", nil)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", test)
	router.HandleFunc("/todo", TodoIndex)
	router.HandleFunc("/todo:{todoId}", TodoShow)

	//	log.Fatal(http.ListenAndServe(":1337", router))

	return router
}
