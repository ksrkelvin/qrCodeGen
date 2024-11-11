package api

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Index is a function that handles the home page
func (p *DiinoAPI) Index(c *gin.Context) {
	fmt.Println("Handling home page")
	// Carrega o template HTML
	tmpl, err := template.ParseFiles("./public/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not load HTML template")
		return
	}
	data := map[string]string{
		"MP_PUBLIC_KEY": os.Getenv("MP_PUBLIC_KEY"),
	}
	// Renderiza o template para o ResponseWriter
	if err := tmpl.Execute(c.Writer, data); err != nil {
		c.String(http.StatusInternalServerError, "Could not render HTML template")
	}
}
