package routes

import (
	"qrCodeGen/internal/api"

	"github.com/gin-gonic/gin"
)

// Router - Router
func Router(app *gin.Engine, diinoAPI *api.DiinoAPI) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	//Index
	app.GET("/", diinoAPI.Home)

	//API
	api := app.Group("/api")

	//Campaign
	api.POST("/generate", diinoAPI.GenQR)
	api.POST("/generateWM", diinoAPI.WatermarkQR)

	//Payment
	payment := api.Group("/payment")
	payment.POST("/preference", diinoAPI.PaymentPreference)
	payment.POST("/confirm", diinoAPI.PaymentConfirm)
	payment.POST("/check", diinoAPI.PaymentCheck)

	return

}
