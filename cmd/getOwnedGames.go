package cmd

import (
	"fmt"
	"log"
	"net/http"
	"steamcli/api"

	"github.com/spf13/cobra"
)

type OwnedGamesResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	GameCount int `json:"game_count"`
	Games     []struct {
		Appid                  int `json:"appid"`
		PlaytimeForever        int `json:"playtime_forever"`
		PlaytimeWindowsForever int `json:"playtime_windows_forever"`
		PlaytimeMacForever     int `json:"playtime_mac_forever"`
		PlaytimeLinuxForever   int `json:"playtime_linux_forever"`
		PlaytimeDeckForever    int `json:"playtime_deck_forever"`
		RtimeLastPlayed        int `json:"rtime_last_played"`
		PlaytimeDisconnected   int `json:"playtime_disconnected"`
		Playtime2Weeks         int `json:"playtime_2weeks,omitempty"`
	} `json:"games"`
}

type AppIDName struct {
	Appid int    `json:"appid"`
	Name  string `json:"name"`
}

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

		ownedGames := new(OwnedGamesResponse)
		url := "http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + api.APIKey + "&steamid=" + api.SteamID + "&format=json"
		err := api.CallSteamAPI(url, http.MethodGet, ownedGames)
		if err != nil {
			log.Fatalf("CallSteamAPI(%v) failed: %v", url, err)
		}
		fmt.Printf("This account has %v games.", ownedGames.Response.GameCount)

		appIDsAndNames := make(map[int]string)

		// THIS IS WORKING BUT IT'S GETTING RATE LIMITED. Steam is returning 429
		// after a few requests, so this won't make it through the entire list.
		// Need to find a way around this. Scrape the App list regularly and
		// write that somewhere?
		for game := range ownedGames.Response.Games {
			fmt.Printf("Looking up %v...\n", ownedGames.Response.Games[game].Appid)
			// This is a public facing API and doesn't require an API key. Perhaps
			// we can get around this by using an internal function instead?
			url := fmt.Sprintf("%s%d", "https://store.steampowered.com/api/appdetails?appids=", ownedGames.Response.Games[game].Appid)
			resp := new(AppIDName)
			err := api.CallSteamAPI(url, http.MethodGet, resp)
			if err != nil {
				log.Fatalf("CallSteamAPI(%v) failed: %v", url, err)
			}
			appIDsAndNames[ownedGames.Response.Games[game].Appid] = resp.Name
		}

		for gameID, gameName := range appIDsAndNames {
			fmt.Printf("\t%v: %v\n", gameID, gameName)
		}
	},
}

func init() {
	rootCmd.AddCommand(getOwnedGamesCmd)
	// TODO(ian): Add a 'count' flag here.
}
