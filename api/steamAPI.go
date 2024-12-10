package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TODO(ian): This can actually call *any* URL that's passed to it. Might need
// to refactor this comment a bit to make that clearer. It's only *supposed* to
// be used to call the Steam API tho.

// CallSteamAPI calls the Steam API based on the given URL string and method. On
// success, it writes to the jsonOutput and returns nil. On an error, it skips
// the json write and returns an error.
func CallSteamAPI(url string, method string, jsonOutput interface{}) error {
	client := &http.Client{
		// Explicitly set timeout so we don't end up with hour long GET calls.
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("could not construct request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error calling steam API with request %v: %v", req, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("steam API returned status: %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(jsonOutput)
	if err != nil {
		return fmt.Errorf("could not unmarshal Steam API response: %v", err)
	}

	return nil
}
