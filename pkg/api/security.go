package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateConnection is a function that validates the connection
func (p *DiinoAPI) ValidateConnection(c *gin.Context) {
	path := c.Request.URL.Path
	ip := c.ClientIP()

	err := p.Diino.Security.UpsertPath(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error validating path"})
		c.Abort()
		return
	}

	if p.Diino.Security.IsProhibitedPath(path) {
		// Salva o IP na lista negra
		p.Diino.Security.BlockIP(ip, path)
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found!"})
		c.Abort()
		return
	}

	c.Next()
}
