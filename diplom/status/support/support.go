package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func StatusSupport(url string) []SupportData {
	result := make([]SupportData, 0)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error() + `: ` + url)
		return []SupportData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []SupportData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return []SupportData{}
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return []SupportData{}
	}

	return result
}
