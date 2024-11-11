package handlefuncs

import (
	"encoding/json"
	"net/http"
	"os"
)

// CheckPaymentRequest Estrutura para capturar o corpo da requisição
type CheckPaymentRequest struct {
	ExternalReference string `json:"externalReference"`
}

// HandleCheckPayment é uma função que verifica o pagamento
func HandleCheckPayment(writer http.ResponseWriter, request *http.Request) {

	var checkRequest CheckPaymentRequest
	err := json.NewDecoder(request.Body).Decode(&checkRequest)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Comparação com a variável de ambiente
	secretKey := os.Getenv("MP_EXTERNAL_REFERENCE")
	if checkRequest.ExternalReference == secretKey {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"status": "approved"}`))
	} else {
		http.Error(writer, "Payment verification failed", http.StatusForbidden)
	}
	return

}
