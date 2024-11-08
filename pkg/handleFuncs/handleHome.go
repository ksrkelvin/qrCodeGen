package handlefuncs

import (
	"fmt"
	"net/http"
	"text/template"
)

// HandleHome is a function that handles the home page
func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Handling home page")
	// Carrega o template HTML
	tmpl, err := template.ParseFiles("./public/index.html")
	if err != nil {
		http.Error(writer, "Could not load HTML template", http.StatusInternalServerError)
		return
	}
	// Renderiza o template para o ResponseWriter
	tmpl.Execute(writer, nil)
}
