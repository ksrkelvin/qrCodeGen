package qrgen

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"

	"github.com/nfnt/resize"
)

// UploadFile is a function that uploads a file
func UploadFile(file multipart.File) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, fmt.Errorf("could not upload file. %v", err)
	}

	return buf.Bytes(), nil
}

// ResizeWatermark redimensiona a marca d'água para o tamanho adequado
func ResizeWatermark(watermarkData *bytes.Reader, width uint) ([]byte, error) {
	watermarkImage, _, err := image.Decode(watermarkData)
	if err != nil {
		return nil, fmt.Errorf("could not decode watermark image: %v", err)
	}

	// Redimensiona a marca d'água
	resizedWatermark := resize.Resize(width, 0, watermarkImage, resize.Lanczos3)

	// Converte a imagem redimensionada para PNG
	var buf bytes.Buffer
	if err := png.Encode(&buf, resizedWatermark); err != nil {
		return nil, fmt.Errorf("could not encode resized watermark: %v", err)
	}

	return buf.Bytes(), nil
}

// applyWatermarkOpacity aplica a opacidade na marca d'água
func applyWatermarkOpacity(watermark image.Image, opacity float64) image.Image {
	bounds := watermark.Bounds()
	rgba := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := watermark.At(x, y)
			r, g, b, a := oldColor.RGBA()
			rgba.Set(x, y, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(float64(a>>8) * opacity), // Aplica a opacidade
			})
		}
	}
	return rgba
}
