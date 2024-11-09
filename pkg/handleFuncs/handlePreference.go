package handlefuncs

import (
	"net/http"
	"qrCodeGen/pkg/mercadopago"
)

// HandlePreference is a function that handles the preference page
func HandlePreference(writer http.ResponseWriter, request *http.Request) {
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
