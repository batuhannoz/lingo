package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "lingo",
		Short: "an app to learn and memorize new words",
		Long:  `translate words into other languages, save as a list and memorize`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root command")
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(translateCmd)
}
