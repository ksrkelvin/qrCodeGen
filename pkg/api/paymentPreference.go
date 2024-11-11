package api

import (
	"net/http"
	"qrCodeGen/pkg/mercadopago"

	"github.com/gin-gonic/gin"
)

// PaymentPreference is a function that handles the preference page
func (p *DiinoAPI) PaymentPreference(c *gin.Context) {
	preference, err := mercadopago.PreferenceProcess()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error processing the payment")
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, preference)
}
