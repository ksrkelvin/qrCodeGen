package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"qrCodeGen/pkg/qrgen"

	"github.com/gin-gonic/gin"
)

// WatermarkQR is a function that handles the QR code generation with watermark
func (p *DiinoAPI) WatermarkQR(c *gin.Context) {
	fmt.Println("Handling QR code generation")
	request.ParseMultipartForm(10 << 20)
	var content string = request.FormValue("content")
	var codeData []byte

	if content == "" {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			"Could not determine the desired QR code content.",
		)
		return
	}

	qrCode := qrgen.SimpleQRCode{Content: content, Size: 256}

	watermarkFile, _, err := request.FormFile("watermark")
	if err != nil && errors.Is(err, http.ErrMissingFile) {
		codeData, err = qrCode.Generate()
		if err != nil {
			writer.WriteHeader(400)
			json.NewEncoder(writer).Encode(
				fmt.Sprintf("Could not generate QR code. %v", err),
			)
			return
		}
		writer.Header().Add("Content-Type", "image/png")
		writer.Write(codeData)
		return
	}

	watermark, err := qrgen.UploadFile(watermarkFile)
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			fmt.Sprint("Could not upload the watermark image.", err),
		)
		return
	}

	codeData, err = qrCode.GenerateWithWatermark(watermark)
	if err != nil {
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(
			fmt.Sprintf(
				"Could not generate QR code with the watermark image. %v", err,
			),
		)
		return
	}

	writer.Header().Set("Content-Type", "image/png")
	writer.Write(codeData)
}
