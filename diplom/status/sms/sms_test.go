package sms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	name     string
	input    string
	expected []SMSData
}{
	{
		name:  "Тест SMS пример из PDF",
		input: "testdata/sms-pdf.data",
		expected: []SMSData{
			{Country: "US", Bandwidth: "36", ResponseTime: "1576", Provider: "Rond"},
			{Country: "BL", Bandwidth: "68", ResponseTime: "1594", Provider: "Kildy"},
		},
	},
	{
		name:  "Тест SMS из симулятора",
		input: "testdata/sms-simulator.data",
		expected: []SMSData{
			{Country: "RU", Bandwidth: "98", ResponseTime: "1576", Provider: "Topolo"},
			{Country: "US", Bandwidth: "97", ResponseTime: "108", Provider: "Rond"},
			{Country: "GB", Bandwidth: "47", ResponseTime: "129", Provider: "Topolo"},
			{Country: "FR", Bandwidth: "18", ResponseTime: "1990", Provider: "Topolo"},
			{Country: "BL", Bandwidth: "16", ResponseTime: "1085", Provider: "Kildy"},
			{Country: "AT", Bandwidth: "17", ResponseTime: "124", Provider: "Topolo"},
			{Country: "BG", Bandwidth: "6", ResponseTime: "1461", Provider: "Rond"},
			{Country: "DK", Bandwidth: "31", ResponseTime: "510", Provider: "Topolo"},
			{Country: "CA", Bandwidth: "35", ResponseTime: "454", Provider: "Rond"},
			{Country: "ES", Bandwidth: "65", ResponseTime: "1231", Provider: "Topolo"},
			{Country: "CH", Bandwidth: "21", ResponseTime: "1674", Provider: "Topolo"},
			{Country: "TR", Bandwidth: "46", ResponseTime: "1790", Provider: "Rond"},
			{Country: "PE", Bandwidth: "46", ResponseTime: "116", Provider: "Topolo"},
			{Country: "NZ", Bandwidth: "60", ResponseTime: "385", Provider: "Kildy"},
			{Country: "MC", Bandwidth: "39", ResponseTime: "718", Provider: "Kildy"},
		},
	},
	{
		name:  "Тест SMS ошибки Country",
		input: "testdata/sms-err-country.data",
		expected: []SMSData{
			{Country: "TR", Bandwidth: "46", ResponseTime: "1790", Provider: "Rond"},
			{Country: "PE", Bandwidth: "46", ResponseTime: "116", Provider: "Topolo"},
			{Country: "MC", Bandwidth: "39", ResponseTime: "718", Provider: "Kildy"},
		},
	},
	{
		name:  "Тест SMS ошибки Bandwidth",
		input: "testdata/sms-err-bandwidth.data",
		expected: []SMSData{
			{Country: "FR", Bandwidth: "18", ResponseTime: "1990", Provider: "Topolo"},
			{Country: "BL", Bandwidth: "16", ResponseTime: "1085", Provider: "Kildy"},
		},
	},
	{
		name:  "Тест SMS ошибки ResponseTime",
		input: "testdata/sms-err-time.data",
		expected: []SMSData{
			{Country: "GB", Bandwidth: "47", ResponseTime: "129", Provider: "Topolo"},
		},
	},
	{
		name:  "Тест SMS ошибки Provider",
		input: "testdata/sms-err-provider.data",
		expected: []SMSData{
			{Country: "RU", Bandwidth: "98", ResponseTime: "1576", Provider: "Topolo"},
			{Country: "US", Bandwidth: "97", ResponseTime: "108", Provider: "Rond"},
			{Country: "BL", Bandwidth: "16", ResponseTime: "1085", Provider: "Kildy"},
		},
	},
	{
		name:     "Тест SMS ошибки All",
		input:    "testdata/sms-err-all.data",
		expected: []SMSData{},
	},
}

func TestStatusSMS(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StatusSMS(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}
