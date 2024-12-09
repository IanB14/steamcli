package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OwnedGamesResponse struct {
	Response struct {
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
	} `json:"response"`
}

// CallSteamAPI uses Steam's Web API to look up the owned games for a given
// account.
func GetOwnedGames(accountID string) (*OwnedGamesResponse, error) {
	// TODO(ian): Use an efficient string builder here
	req := "http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + APIKey + "&steamid=" + accountID + "&format=json"
	response, err := http.Get(req)
	if err != nil {
		return nil, fmt.Errorf("GetOwnedGames() for accountID %v failed: %v", accountID, err)
	}
	// TODO(ian): Consider a !200 check here
	defer response.Body.Close()

	responseBody := new(OwnedGamesResponse)
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal Steam API response: %v", err)
	}

	return responseBody, nil
}
