package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/jonioliveira/tlt/internal/credentials"
	"github.com/jonioliveira/tlt/internal/translate"
	"github.com/spf13/cobra"
)

func GetTranslateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "translate",
		Short: "translate sentence",
		Long:  "translate sentence",
		Run: func(cmd *cobra.Command, args []string) {
			context := context.Background()
			credentials, err := credentials.GetEnvCreds("TRANSLATION_API_CREDS")
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
			translated, err := translate.Do(context, credentials, strings.Join(args, " "), "PT", "EN")
			if err != nil {
				fmt.Printf("Error: %s", err)
			}

			fmt.Printf("%s", translated)
		},
	}
}
