package main

import (
	"fmt"
	"os"

	"github.com/jonioliveira/tlt/internal/cmd"
)

func main() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(cmd.GetVersionCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// ctx := context.Background()
	// creds, err := getCreds()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// result, err := translate.Do(ctx, creds, "Ola", "PT", "EN")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(result)
}

// func getCreds() ([]byte, error) {
// 	encodedCreds := os.Getenv("TRANSLATION_API_CREDS")

// 	creds, err := base64.StdEncoding.DecodeString(encodedCreds)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return creds, nil
// }
