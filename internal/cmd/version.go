package cmd

import (
	"fmt"

	"github.com/jonioliveira/tlt/internal/version"
	"github.com/spf13/cobra"
)

func GetVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display tlt version",
		Long:  "Display tlt version",
		Run: func(cmd *cobra.Command, args []string) {
			version, err := version.GetVersion()
			if err != nil {
				fmt.Printf("%v", err)
			}

			fmt.Printf("TLT version : %s", version)
		},
	}
}
