package api

import (
	"fmt"
	"net/http"
	"qrCodeGen/pkg/qrgen"

	"github.com/gin-gonic/gin"
)

// GenQR is a function that handles the QR code generation
func (p *DiinoAPI) GenQR(c *gin.Context) {
	fmt.Println("Handling QR code generation")
	var content = c.Request.FormValue("content")
	var codeData []byte

	if content == "" {
		c.String(http.StatusBadRequest, `{"message": "content is required"}`)
		return
	}
	qrCode := qrgen.SimpleQRCode{Content: content, Size: 256}
	codeData, err := qrCode.Generate()
	if err != nil {
		c.String(http.StatusInternalServerError, `{"message": "error generating qr code"}`)
		return
	}
	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", codeData)

}
