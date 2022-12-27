package sms

import (
	"bufio"
	"fmt"
	"os"
	"skillbox/diplom/status/check"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func SMSChangeCodeToCountry(data []SMSData) []SMSData {
	res := make([]SMSData, 0)

	for _, elem := range data {
		country := check.GetCountryForCode(elem.Country)
		elem.Country = country
		res = append(res, elem)
	}

	return res
}

func ParseSMSData(line string) (SMSData, bool) {
	data := strings.Split(line, ";")

	if len(data) != 4 {
		return SMSData{}, false
	}

	Country := data[0]
	if !check.IsCountry(Country) {
		return SMSData{}, false
	}

	Bandwidth := data[1]
	if !check.IsBandwidth(Bandwidth) {
		return SMSData{}, false
	}

	ResponseTime := data[2]
	if !check.IsResponseTime(ResponseTime) {
		return SMSData{}, false
	}

	Provider := data[3]
	if !check.IsProviderSMSandMMS(Provider) {
		return SMSData{}, false
	}

	elem := SMSData{
		Country:      Country,
		Bandwidth:    Bandwidth,
		ResponseTime: ResponseTime,
		Provider:     Provider,
	}

	return elem, true
}

func StatusSMS(csvFile string) []SMSData {
	result := make([]SMSData, 0)

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return []SMSData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		elem, ok := ParseSMSData(line)

		if ok {
			result = append(result, elem)
		}
	}

	return result
}
