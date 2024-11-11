package qrgen

import "github.com/skip2/go-qrcode"

// Generate generates a QR code
func (p *SimpleQRCode) Generate() (data []byte, err error) {
	data, err = qrcode.Encode(p.Content, qrcode.Medium, p.Size)
	if err != nil {
		return nil, err
	}
	return
}
