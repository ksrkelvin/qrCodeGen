package qrgen

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
)

// GenerateWithWatermark generates a QR code with a watermark
func (code *SimpleQRCode) GenerateWithWatermark(watermark []byte) ([]byte, error) {
	qrCode, err := code.Generate()
	if err != nil {
		return nil, err
	}

	qrCode, err = code.addWatermark(qrCode, watermark)
	if err != nil {
		return nil, fmt.Errorf("could not add watermark to QR code: %v", err)
	}

	return qrCode, nil
}

// addWatermark adiciona a marca d'água ao QR code
func (code *SimpleQRCode) addWatermark(qrCode []byte, watermarkData []byte) ([]byte, error) {
	// Decodifica o QR code gerado
	qrCodeData, err := png.Decode(bytes.NewBuffer(qrCode))
	if err != nil {
		return nil, fmt.Errorf("could not decode QR code: %v", err)
	}

	// Redimensiona a marca d'água
	watermarkWidth := uint(float64(qrCodeData.Bounds().Dx()) * 0.2) // Marca d'água com 20% da largura do QR code
	watermark, err := ResizeWatermark(bytes.NewReader(watermarkData), watermarkWidth)
	if err != nil {
		return nil, fmt.Errorf("could not resize the watermark image: %v", err)
	}

	watermarkImage, err := png.Decode(bytes.NewBuffer(watermark))
	if err != nil {
		return nil, fmt.Errorf("could not decode watermark: %v", err)
	}

	// Calcula o ponto de centralização da marca d'água
	qrCodeBounds := qrCodeData.Bounds()
	watermarkBounds := watermarkImage.Bounds()
	watermarkX := (qrCodeBounds.Dx() - watermarkBounds.Dx()) / 2
	watermarkY := (qrCodeBounds.Dy() - watermarkBounds.Dy()) / 2
	watermarkOffset := image.Pt(watermarkX, watermarkY)

	// Cria uma nova imagem para o QR code com a marca d'água
	m := image.NewRGBA(qrCodeBounds)

	// Desenha o QR code primeiro
	draw.Draw(m, qrCodeBounds, qrCodeData, image.Point{}, draw.Src)

	// Aplica a marca d'água com opacidade
	watermarkWithOpacity := applyWatermarkOpacity(watermarkImage, 8) // 20% de opacidade
	draw.Draw(
		m,
		watermarkWithOpacity.Bounds().Add(watermarkOffset),
		watermarkWithOpacity,
		image.Point{},
		draw.Over,
	)

	// Cria o QR code final com a marca d'água
	watermarkedQRCode := bytes.NewBuffer(nil)
	png.Encode(watermarkedQRCode, m)

	return watermarkedQRCode.Bytes(), nil
}
