package cmd

import (

	"github.com/batuhannoz/lingo/data"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	translateCmd = &cobra.Command{
		Use:   "translate",
		Short: "an app to learn and memorize new words",
		Long:  `translate words into other languages, save as a list and memorize`,
		Run: func(cmd *cobra.Command, args []string) {
			data.CreateTable()
		},
	}
)

func init() {

}
