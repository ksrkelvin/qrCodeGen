package main

import (
	"fmt"
	"net/http"
	handlefuncs "qrCodeGen/pkg/handleFuncs"
)

func main() {
	http.HandleFunc("/", handlefuncs.HandleHome)

	http.HandleFunc("/generate", handlefuncs.HandleQR)
	http.HandleFunc("/generateWM", handlefuncs.HandleQRWM)
	http.HandleFunc("/preference", handlefuncs.HandlePreference)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port 8080")
}
