package qrgen

import "github.com/skip2/go-qrcode"

// SimpleQRCode is a struct that represents a simple QR code
type SimpleQRCode struct {
	Content string
	Size    int
}

// Generate generates a QR code
func (p *SimpleQRCode) Generate() (data []byte, err error) {
	data, err = qrcode.Encode(p.Content, qrcode.Medium, p.Size)
	if err != nil {
		return nil, err
	}
	return
}
