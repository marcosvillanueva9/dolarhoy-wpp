package send

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Run(mensaje string) {
	apiUrl := "https://api.twilio.com/2010-04-01/Accounts/ACd9518e92486ab2547dc1f3b7525e2e40/Messages.json"
    data := url.Values{}
    data.Add("To", "whatsapp:+5492612515909")
    data.Add("From", "whatsapp:+14155238886")
    data.Add("Body", mensaje)

    encodedData := data.Encode()

    client := &http.Client{}
    r, _ := http.NewRequest("POST", apiUrl, strings.NewReader(encodedData))
	r.Header.Set("Authorization", "Basic QUNkOTUxOGU5MjQ4NmFiMjU0N2RjMWYzYjc1MjVlMmU0MDpkNTM5NTFlMGUyZGUxZmU4OThmYzYzOGVjZDhiYWYxZA==")
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
    fmt.Println(resp.Status)
	fmt.Println(resp.Body)
}