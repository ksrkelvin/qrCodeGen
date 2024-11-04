package main

import (
	"net/http"
	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var size, content = request.FormValue("size"), request.FormValue("content")
	var codeData []byte

	writer.Header().Set("Content-Type", "application/json")

	if content == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"message": "content is required"}`))
		return
	}

	qrCodeSize, err := strconv.Atoi(size)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"message": "size is required"}`))
		return
	}

	qrCode := simpleQRCode{Content: content, Size: qrCodeSize}
	codeData, err = qrCode.Generate()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"message": "error generating qr code"}`))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	writer.Write(codeData)

}
func main() {
	http.HandleFunc("/generate", handleRequest)
	http.ListenAndServe(":8080", nil)
}

type simpleQRCode struct {
	Content string
	Size    int
}

func (p *simpleQRCode) Generate() (data []byte, err error) {
	data, err = qrcode.Encode(p.Content, qrcode.Medium, p.Size)
	if err != nil {
		return nil, err
	}
	return
}
