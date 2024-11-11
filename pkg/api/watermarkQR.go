package api

import (
	"errors"
	"fmt"
	"net/http"
	"qrCodeGen/pkg/qrgen"

	"github.com/gin-gonic/gin"
)

// WatermarkQR is a function that handles the QR code generation with watermark
func (p *DiinoAPI) WatermarkQR(c *gin.Context) {
	fmt.Println("Handling QR code generation")
	var content string = c.Request.FormValue("content")
	var codeData []byte

	if content == "" {
		c.String(http.StatusBadRequest, "Could not determine the desired QR code content.")
		return
	}

	qrCode := qrgen.SimpleQRCode{Content: content, Size: 256}

	watermarkFile, _, err := c.Request.FormFile("watermark")
	if err != nil && errors.Is(err, http.ErrMissingFile) {
		codeData, err = qrCode.Generate()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Could not generate QR code. %v", err))
			return
		}
		c.Data(http.StatusBadRequest, "image/png", codeData)
		return
	}

	watermark, err := qrgen.UploadFile(watermarkFile)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint("Could not upload the watermark image.", err))
		return
	}

	codeData, err = qrCode.GenerateWithWatermark(watermark)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Could not generate QR code with the watermark image. %v", err))
		return
	}

	c.Data(http.StatusBadRequest, "image/png", codeData)
}
