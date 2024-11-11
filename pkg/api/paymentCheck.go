package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CheckPaymentRequest Estrutura para capturar o corpo da requisição
type CheckPaymentRequest struct {
	ExternalReference string `json:"externalReference"`
}

// PaymentCheck é uma função que verifica o pagamento
func (p *DiinoAPI) PaymentCheck(c *gin.Context) {

	var checkRequest CheckPaymentRequest
	err := json.NewDecoder(c.Request.Body).Decode(&checkRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Comparação com a variável de ambiente
	secretKey := os.Getenv("MP_EXTERNAL_REFERENCE")
	if checkRequest.ExternalReference == secretKey {
		c.String(http.StatusOK, `{"status": "approved"}`)
	} else {
		c.String(http.StatusForbidden, "Payment verification failed")
	}
	return

}
