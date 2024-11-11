package handlefuncs

import (
	"fmt"
	"net/http"
)

// HandlePaymentConfirm is a function that handles the Payment confirmation
func HandlePaymentConfirm(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Println(request)
}
