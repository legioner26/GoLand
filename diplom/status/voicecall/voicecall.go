package voicecall

import (
	"bufio"
	"fmt"
	"os"
	"skillbox/diplom/status/check"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_call_time"`
}

func ParseVoiceCallData(line string) (VoiceCallData, bool) {
	data := strings.Split(line, ";")

	if len(data) != 8 {
		return VoiceCallData{}, false
	}

	Country := data[0]
	if !check.IsCountry(Country) {
		return VoiceCallData{}, false
	}

	Bandwidth := data[1]
	if !check.IsBandwidth(Bandwidth) {
		return VoiceCallData{}, false
	}

	ResponseTime := data[2]
	if !check.IsResponseTime(ResponseTime) {
		return VoiceCallData{}, false
	}

	Provider := data[3]
	if !check.IsProviderVoiceCall(Provider) {
		return VoiceCallData{}, false
	}

	ConnectionStability, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return VoiceCallData{}, false
	}

	TTFB, err := strconv.Atoi(data[5])
	if err != nil {
		return VoiceCallData{}, false
	}

	VoicePurity, err := strconv.Atoi(data[6])
	if err != nil {
		return VoiceCallData{}, false
	}

	MedianOfCallsTime, err := strconv.Atoi(data[7])
	if err != nil {
		return VoiceCallData{}, false
	}

	elem := VoiceCallData{
		Country:             Country,
		Bandwidth:           Bandwidth,
		ResponseTime:        ResponseTime,
		Provider:            Provider,
		ConnectionStability: float32(ConnectionStability),
		TTFB:                TTFB,
		VoicePurity:         VoicePurity,
		MedianOfCallsTime:   MedianOfCallsTime,
	}

	return elem, true
}

func StatusVoiceCall(csvFile string) []VoiceCallData {
	result := make([]VoiceCallData, 0)

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return []VoiceCallData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		elem, ok := ParseVoiceCallData(line)

		if ok {
			result = append(result, elem)
		}
	}

	return result
}
