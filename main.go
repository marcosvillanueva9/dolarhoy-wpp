package main

import (
	"fmt"
	"strings"

	"github.com/marcosvillanueva9/dolarhoy-wpp/scrapper"
)

func main() {
		valordolar := scrapper.Run()

		fmt.Println(formatMessage(valordolar))
}

func formatMessage(valordolar string) string{

	vals := strings.Split(valordolar, "$")

	if len(vals) < 3 {
		return "error en lo que vino del scrapper: " + valordolar 
	}

	message := fmt.Sprintf("Dolar compra: $%s | Dolar venta: $%s", vals[1], vals[2])
	return message
}