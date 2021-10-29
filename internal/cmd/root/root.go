package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tlt",
		Short: "Translation CLI app for fast translations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("TLT CLI")
		},
	}
}
