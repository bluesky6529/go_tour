package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本",
	Long:  `版本`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vocabulary v0.0.1")
	},
}
