package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/marcosvillanueva9/dolarhoy-wpp/scrapper"
	"github.com/marcosvillanueva9/dolarhoy-wpp/send"
	"github.com/robfig/cron/v3"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		// Get Dolar Value
		valordolar := scrapper.Run()

		c.JSON(http.StatusOK, gin.H{"data": formatMessage(valordolar)})
	})

	r.GET("/send", func(c *gin.Context) {

		// Get Dolar Value
		valordolar := scrapper.Run()

		mensaje := formatMessage(valordolar)

		send.Run(mensaje)

		c.JSON(http.StatusOK, gin.H{"data": "sent: " + mensaje})
	})

	go func() {
		c := cron.New()
		c.AddFunc("@every 1d", func() { 

			// Get Dolar Value
			valordolar := scrapper.Run()

			mensaje := formatMessage(valordolar)

			send.Run(mensaje)
		})

		// Start cron with one scheduled job
		fmt.Println("Start cron")
		c.Start()
	}()

	r.Run()
}

func formatMessage(valordolar string) string{

	vals := strings.Split(valordolar, "$")

	if len(vals) < 3 {
		return "error en lo que vino del scrapper: " + valordolar 
	}

	message := fmt.Sprintf("Dolar compra: $%s | Dolar venta: $%s", vals[1], vals[2])
	return message
}