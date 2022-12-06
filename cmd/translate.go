package cmd

import (
	"fmt"

	"github.com/batuhannoz/lingo/data"
	"github.com/batuhannoz/lingo/translate"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	Save bool
	translateCmd = &cobra.Command{
		Use:   "translate",
		Short: "an app to learn and memorize new words",
		Long:  `translate words into other languages, save as a list and memorize`,
		Run: func(cmd *cobra.Command, args []string) {
			text, _ := translate.TranslateText(args[1], args[2])
			fmt.Println(text)
			fmt.Println("save?", Save)
			data.CreateTable()
		},
	}
)

func init() {
	translateCmd.Flags().BoolVarP(&Save, "save", "s", false, "used to save in the translation list")
}
