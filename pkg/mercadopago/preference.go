package mercadopago

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

// PreferenceProcess is a function that processes a payment
func PreferenceProcess() (data string, err error) {
	cfg, err := config.New(os.Getenv("MP_ACCESS_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	urlApp := os.Getenv("URL_APP")
	fmt.Println(urlApp)
	client := preference.NewClient(cfg)

	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				Title:     "QR Code Generator",
				Quantity:  1,
				UnitPrice: 1.00,
			},
		},
		BackURLs: &preference.BackURLsRequest{
			Success: urlApp + "/success",
			Pending: urlApp + "/pending",
			Failure: urlApp + "/failure",
		},
		AutoReturn: "approved",
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonRaw, err := json.Marshal(resource)

	return string(jsonRaw), err
}
