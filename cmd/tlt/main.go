package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/jonioliveira/tlt/internal/translate"
)

func main() {
	ctx := context.Background()
	creds, err := getCreds()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := translate.Do(ctx, creds, "Ola", "PT", "EN")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func getCreds() ([]byte, error) {
	encodedCreds := os.Getenv("TRANSLATION_API_CREDS")

	creds, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
