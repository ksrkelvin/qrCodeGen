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

// addWatermark adds a watermark to a QR code
func (code *SimpleQRCode) addWatermark(qrCode []byte, watermarkData []byte) ([]byte, error) {
	qrCodeData, err := png.Decode(bytes.NewBuffer(qrCode))
	if err != nil {
		return nil, fmt.Errorf("could not decode QR code: %v", err)
	}

	watermarkWidth := uint(float64(qrCodeData.Bounds().Dx()) * 0.25)
	watermark, err := ResizeWatermark(bytes.NewBuffer(watermarkData), watermarkWidth)
	if err != nil {
		return nil, fmt.Errorf("could not resize the watermark image: %v", err)
	}

	watermarkImage, err := png.Decode(bytes.NewBuffer(watermark))
	if err != nil {
		return nil, fmt.Errorf("could not decode watermark: %v", err)
	}

	var halfQrCodeWidth, halfWatermarkWidth int = qrCodeData.Bounds().Dx() / 2, watermarkImage.Bounds().Dx() / 2
	offset := image.Pt(
		halfQrCodeWidth-halfWatermarkWidth,
		halfQrCodeWidth-halfWatermarkWidth,
	)

	watermarkImageBounds := qrCodeData.Bounds()
	m := image.NewRGBA(watermarkImageBounds)

	draw.Draw(m, watermarkImageBounds, qrCodeData, image.Point{}, draw.Src)
	draw.Draw(
		m,
		watermarkImage.Bounds().Add(offset),
		watermarkImage,
		image.Point{},
		draw.Over,
	)

	watermarkedQRCode := bytes.NewBuffer(nil)
	png.Encode(watermarkedQRCode, m)

	return watermarkedQRCode.Bytes(), nil
}
