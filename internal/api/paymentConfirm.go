package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlePaymentConfirm is a function that handles the Payment confirmation
func (p *DiinoAPI) PaymentConfirm(c *gin.Context) {
	writer.WriteHeader(http.StatusOK)
	fmt.Println(request)
}
