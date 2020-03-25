package main

// imports principais
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// resposta para o front
		fmt.Fprintf(w, "bem vindo")
	})
	// iniciar servidor na porta localhost:1337
	println("Servidor iniciado na porta 1337")
	http.ListenAndServe(":1337", nil)
}
