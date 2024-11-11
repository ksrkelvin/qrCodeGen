package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PaymentConfirm is a function that handles the Payment confirmation
func (p *DiinoAPI) PaymentConfirm(c *gin.Context) {
	body := c.Request.Body

	bodyRaw, err := json.Marshal(body)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request payload")
		return

	}

	c.String(http.StatusOK, string(bodyRaw))
}
