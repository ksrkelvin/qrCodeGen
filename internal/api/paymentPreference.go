package api

import (
	"net/http"
	"qrCodeGen/pkg/mercadopago"

	"github.com/gin-gonic/gin"
)

// Preference is a function that handles the preference page
func (p *DiinoAPI) PaymentPreference(c *gin.Context) {
	preference, err := mercadopago.PreferenceProcess()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error processing the payment"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(preference))
}
