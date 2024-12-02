package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseViaCep struct {
	Address string `json:"logradouro"`
	City    string `json:"localidade"`
	State   string `json:"uf"`
	API     string
}

type ResponseBrasilApi struct {
	Address string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	API     string
}

func FasterAPI(cep string) error {

	cepR1 := make(chan ResponseBrasilApi)
	cepR2 := make(chan ResponseViaCep)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		url1 := "https://brasilapi.com.br/api/cep/v1/" + cep

		req, _ := http.NewRequestWithContext(ctx, "GET", url1, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		var data ResponseBrasilApi
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return
		}
		data.API = url1
		cepR1 <- data
	}()

	go func() {

		url2 := "http://viacep.com.br/ws/" + cep + "/json/"

		req, _ := http.NewRequestWithContext(ctx, "GET", url2, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		var data ResponseViaCep
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return
		}
		data.API = url2
		cepR2 <- data
	}()

	select {
	case msg := <-cepR1:
		fmt.Printf("Received from %s\n", "Brasil API")
		fmt.Printf("Endereço: %s Cidade: %s Estado: %s\n", msg.Address, msg.City, msg.State)
		fmt.Printf("The url that send the request was: %s", msg.API)
	case msg := <-cepR2:
		fmt.Printf("Received from %s\n", "Via CEP")
		fmt.Printf("Endereço: %s Cidade: %s Estado: %s\n", msg.Address, msg.City, msg.State)
		fmt.Printf("The url that send the request was: %s", msg.API)
	case <-time.After(time.Second * 1):
		println("timeout")
	}

	return nil
}
