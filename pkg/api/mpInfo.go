package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// MPInfo is a function that handles the MercadoPago public key
func (p *DiinoAPI) MPInfo(c *gin.Context) {

	data := gin.H{
		"MP_PUBLIC_KEY": os.Getenv("MP_PUBLIC_KEY"),
	}

	c.JSON(http.StatusOK, data)
}
