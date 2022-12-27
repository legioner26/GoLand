package email

import (
	"bufio"
	"fmt"
	"os"
	"skillbox/diplom/status/check"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func Get3MinDeliveryTimeByCountry(data []EmailData, code string) []EmailData {
	result := make([]EmailData, 0)
	for _, elem := range data {
		if elem.Country == code {
			result = append(result, elem)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].DeliveryTime < result[j].DeliveryTime
	})

	if len(result) < 3 {
		return result
	}

	return result[:3]
}

func Get3MaxDeliveryTimeByCountry(data []EmailData, code string) []EmailData {
	result := make([]EmailData, 0)
	for _, elem := range data {
		if elem.Country == code {
			result = append(result, elem)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].DeliveryTime > result[j].DeliveryTime
	})

	if len(result) < 3 {
		return result
	}

	return result[:3]
}

func ParseEmailData(line string) (EmailData, bool) {
	data := strings.Split(line, ";")

	if len(data) != 3 {
		return EmailData{}, false
	}

	Country := data[0]
	if !check.IsCountry(Country) {
		return EmailData{}, false
	}

	Provider := data[1]
	if !check.IsProviderEmail(Provider) {
		return EmailData{}, false
	}

	DeliveryTime, err := strconv.Atoi(data[2])
	if err != nil {
		return EmailData{}, false
	}

	elem := EmailData{
		Country:      Country,
		Provider:     Provider,
		DeliveryTime: DeliveryTime,
	}

	return elem, true
}

func StatusEmail(csvFile string) []EmailData {
	result := make([]EmailData, 0)

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return []EmailData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		elem, ok := ParseEmailData(line)

		if ok {
			result = append(result, elem)
		}
	}

	return result
}
