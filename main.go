package main

import (
	"fmt"
	"net/http"
	handlefuncs "qrCodeGen/pkg/handleFuncs"
)

func main() {
	http.HandleFunc("/", handlefuncs.HandleHome)

	http.HandleFunc("/generate", handlefuncs.HandleQR)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port 8080")
}
