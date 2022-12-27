package incident

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}

func StatusIncident(url string) []IncidentData {
	result := make([]IncidentData, 0)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error() + `: ` + url)
		return []IncidentData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []IncidentData{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return []IncidentData{}
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return []IncidentData{}
	}

	return result
}
