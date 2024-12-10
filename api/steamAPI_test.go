package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TODO(ian): This struct is copied from `getOwnedGames`. This should be
// refactored out into a separate package rather than duplicating it here.
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

func TestCallSteamAPI(t *testing.T) {
	tests := []struct {
		name    string
		method  string
		status  int
		want    interface{}
		wantErr bool
	}{
		{
			name:   "valid_request",
			method: http.MethodGet,
			status: http.StatusOK,
			want: &OwnedGamesResponse{
				Response{
					GameCount: 10,
				},
			},
		},
		{
			name:    "invalid_request",
			method:  http.MethodGet,
			status:  http.StatusBadRequest,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
				writer.WriteHeader(test.status)
				output, err := json.Marshal(test.want)
				if err != nil {
					t.Errorf("Could not Marshal %v to JSON: %v", test.want, err)
				}
				writer.Write(output)
			}))
			defer server.Close()

			var got = new(OwnedGamesResponse)
			err := CallSteamAPI(server.URL, http.MethodGet, got)
			if test.wantErr {
				if err == nil {
					t.Errorf("no error occured - expected an error")
				}
			} else {
				if err != nil {
					t.Errorf("an unexpected error occurred - got %v, want nil", err)
				}
			}

			if test.want != nil {
				if diff := cmp.Diff(got, test.want); diff != "" {
					t.Errorf("Got %v, want %v", got, test.want)
				}

			}
		})
	}
}
