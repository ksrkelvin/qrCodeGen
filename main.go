package main

import (
	"fmt"
	"net/http"
	"os"
	handlefuncs "qrCodeGen/pkg/handleFuncs"
	"qrCodeGen/pkg/tools"
	"time"
)

func main() {
	tools.CreateHash(os.Getenv("MP_EXTERNAL_REFERENCE") + time.Now().String())
	fmt.Println("Secret Key: ", tools.GetSecretKey())
	http.HandleFunc("/", handlefuncs.HandleHome)

	http.HandleFunc("/generate", handlefuncs.HandleQR)
	http.HandleFunc("/generateWM", handlefuncs.HandleQRWM)
	http.HandleFunc("/preference", handlefuncs.HandlePreference)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port 8080")
}
