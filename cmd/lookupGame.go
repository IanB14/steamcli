package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lookupGameCmd represents the lookupGame command
var lookupGameCmd = &cobra.Command{
	Use:   "lookupGame",
	Short: "Look up game in the Steam store based on a given string.",
	Long:  `TODO: longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lookupGame called")
	},
}

func init() {
	rootCmd.AddCommand(lookupGameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lookupGameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lookupGameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
