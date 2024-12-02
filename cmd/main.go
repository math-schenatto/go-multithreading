package main

import (
	"log"

	apiclient "github.com/math-sche/go-multithreading/internal/apiClient"
)

func main() {
	cep := "01153000"

	err := apiclient.FasterAPI(cep)
	if err != nil {
		log.Println(err)
	}

}
