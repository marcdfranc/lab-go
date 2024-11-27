package router

/*
import (
	"fmt"
	"net/http"
)

func main() {
	router := NewRouter()

	// Definindo uma rota com variável na URL
	router.HandleFunc("/article/{id}", articleHandler)

	// Definindo uma rota com parâmetros de consulta
	router.HandleFunc("/search", searchHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

// Manipulador para a rota com variável na URL
func articleHandler(w http.ResponseWriter, r *http.Request) {
	id := Param(r, "id")
	fmt.Fprintf(w, "Você está vendo o artigo com ID: %s\n", id)
}

// Manipulador para a rota com parâmetros de consulta
func searchHandler(w http.ResponseWriter, r *http.Request) {
	paramA := Query(r, "parama")
	paramB := Query(r, "paramb")
	fmt.Fprintf(w, "Resultados da busca para parama=%s e paramb=%s\n", paramA, paramB)
}
*/
