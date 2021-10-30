package main

import (
	"fmt"
	"os"

	"github.com/jonioliveira/tlt/internal/cmd"
)

func main() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(cmd.GetVersionCmd())
	rootCmd.AddCommand(cmd.GetTranslateCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// ctx := context.Background()
}
