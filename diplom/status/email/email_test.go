package email

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	name     string
	input    string
	expected []EmailData
}{
	{
		name:  "Тест Email пример из PDF",
		input: "testdata/email-pdf.data",
		expected: []EmailData{
			{Country: "AT", Provider: "Hotmail", DeliveryTime: 487},
		},
	},
	{
		name:  "Тест Email из симулятора",
		input: "testdata/email-simulator.data",
		expected: []EmailData{
			{Country: "RU", Provider: "Gmail", DeliveryTime: 106},
			{Country: "RU", Provider: "Yahoo", DeliveryTime: 515},
			{Country: "RU", Provider: "Hotmail", DeliveryTime: 119},
			{Country: "RU", Provider: "MSN", DeliveryTime: 316},
			{Country: "RU", Provider: "Orange", DeliveryTime: 415},
			{Country: "RU", Provider: "Comcast", DeliveryTime: 211},
			{Country: "RU", Provider: "AOL", DeliveryTime: 204},
			{Country: "RU", Provider: "Live", DeliveryTime: 260},
			{Country: "RU", Provider: "RediffMail", DeliveryTime: 397},
			{Country: "RU", Provider: "GMX", DeliveryTime: 457},
			{Country: "RU", Provider: "Protonmail", DeliveryTime: 62},
			{Country: "RU", Provider: "Yandex", DeliveryTime: 504},
			{Country: "RU", Provider: "Mail.ru", DeliveryTime: 392},
		},
	},
	{
		name:     "Тест Email ошибки All",
		input:    "testdata/email-err-all.data",
		expected: []EmailData{},
	},
}

func TestStatusEmail(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StatusEmail(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}
