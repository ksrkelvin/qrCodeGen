package handlefuncs

import (
	"fmt"
	"net/http"
	"qrCodeGen/pkg/qrgen"
)

// HandleQR is a function that handles the QR code generation
func HandleQR(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Handling QR code generation")
	request.ParseMultipartForm(10 << 20)
	var content = request.FormValue("content")
	var codeData []byte

	writer.Header().Set("Content-Type", "application/json")

	if content == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"message": "content is required"}`))
		return
	}
	qrCode := qrgen.SimpleQRCode{Content: content, Size: 256}
	codeData, err := qrCode.Generate()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"message": "error generating qr code"}`))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	writer.Write(codeData)

}
