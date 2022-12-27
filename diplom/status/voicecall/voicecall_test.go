package voicecall

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	name     string
	input    string
	expected []VoiceCallData
}{
	{
		name:  "Тест VoiceCall пример из PDF",
		input: "testdata/voicecall-pdf.data",
		expected: []VoiceCallData{
			{Country: "BG", Bandwidth: "40", ResponseTime: "609", Provider: "E-Voice",
				ConnectionStability: 0.86, TTFB: 160, VoicePurity: 36, MedianOfCallsTime: 5},
			{Country: "DK", Bandwidth: "11", ResponseTime: "743", Provider: "JustPhone",
				ConnectionStability: 0.67, TTFB: 82, VoicePurity: 74, MedianOfCallsTime: 41},
		},
	},
	{
		name:  "Тест VoiceCall из симулятора",
		input: "testdata/voicecall-simulator.data",
		expected: []VoiceCallData{
			{Country: "RU", Bandwidth: "51", ResponseTime: "1772", Provider: "TransparentCalls",
				ConnectionStability: 0.97, TTFB: 402, VoicePurity: 17, MedianOfCallsTime: 10},
			{Country: "US", Bandwidth: "78", ResponseTime: "669", Provider: "E-Voice",
				ConnectionStability: 0.78, TTFB: 62, VoicePurity: 70, MedianOfCallsTime: 31},
			{Country: "DK", Bandwidth: "15", ResponseTime: "626", Provider: "JustPhone",
				ConnectionStability: 0.73, TTFB: 833, VoicePurity: 78, MedianOfCallsTime: 27},
		},
	},
	{
		name:     "Тест VoiceCall ошибки All",
		input:    "testdata/voicecall-err-all.data",
		expected: []VoiceCallData{},
	},
}

func TestStatusVoiceCall(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StatusVoiceCall(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}
