package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index is a function that handles the home page
func (p *DiinoAPI) Index(c *gin.Context) {

	fmt.Println("Handling home page")
	// Data passed to the template

	// Render the HTML template
	c.HTML(http.StatusOK, "index.html", nil)
}
