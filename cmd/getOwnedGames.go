package cmd

import (
	"fmt"
	"log"
	"steamcli/api"

	"github.com/spf13/cobra"
)

// TODO(ian): getOwnedGames currently returns a *count* rather than an actual
// list of games.
var getOwnedGamesCmd = &cobra.Command{
	Use:     "getOwnedGames",
	Aliases: []string{"mygames"},
	Short:   "Lookup all games owned by a particular account.",
	Long:    `TODO(ian): Write a proper description here.`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			log.Fatalf("an account ID must be specified.")
		}

		ownedGamesResponse, err := api.GetOwnedGames(args[0])
		if err != nil {
			log.Fatalf("Failed to look up owned games: %v", err)
		}

		fmt.Printf("This account has %v games.", ownedGamesResponse.Response.GameCount)

	},
}

func init() {
	rootCmd.AddCommand(getOwnedGamesCmd)
}
