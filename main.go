package main

<<<<<<< HEAD
// imports principais
=======
>>>>>>> 3815321d7a6648f68ea23155243087f6e819793f
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		// resposta para o front
=======
>>>>>>> 3815321d7a6648f68ea23155243087f6e819793f
		fmt.Fprintf(w, "bem vindo")
	})
	// iniciar servidor na porta localhost:1337
	println("Servidor iniciado na porta 1337")
<<<<<<< HEAD

=======
>>>>>>> 3815321d7a6648f68ea23155243087f6e819793f
	http.ListenAndServe(":1337", nil)
}
