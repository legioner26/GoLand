package billing

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var tests = []struct {
	name     string
	input    string
	expected BillingData
}{
	{
		name:  "Тест Billing пример из PDF",
		input: "testdata/billing-pdf.data",
		expected: BillingData{
			CreateCustomer: true,
			Purchase:       true,
			Payout:         false,
			Recurring:      false,
			FraudControl:   true,
			CheckoutPage:   true,
		},
	},
	{
		name:  "Тест Billing из симулятора",
		input: "testdata/billing-simulator.data",
		expected: BillingData{
			CreateCustomer: true,
			Purchase:       true,
			Payout:         true,
			Recurring:      true,
			FraudControl:   false,
			CheckoutPage:   false,
		},
	},
	{
		name:  "Тест Billing 1",
		input: "testdata/billing-1.data",
		expected: BillingData{
			CreateCustomer: true,
			Purchase:       false,
			Payout:         false,
			Recurring:      false,
			FraudControl:   false,
			CheckoutPage:   false,
		},
	},
	{
		name:  "Тест Billing 2",
		input: "testdata/billing-2.data",
		expected: BillingData{
			CreateCustomer: false,
			Purchase:       true,
			Payout:         false,
			Recurring:      false,
			FraudControl:   false,
			CheckoutPage:   false,
		},
	},
	{
		name:  "Тест Billing 3",
		input: "testdata/billing-3.data",
		expected: BillingData{
			CreateCustomer: false,
			Purchase:       false,
			Payout:         true,
			Recurring:      false,
			FraudControl:   false,
			CheckoutPage:   false,
		},
	},
	{
		name:  "Тест Billing 4",
		input: "testdata/billing-4.data",
		expected: BillingData{
			CreateCustomer: false,
			Purchase:       false,
			Payout:         false,
			Recurring:      true,
			FraudControl:   false,
			CheckoutPage:   false,
		},
	},
	{
		name:  "Тест Billing 5",
		input: "testdata/billing-5.data",
		expected: BillingData{
			CreateCustomer: false,
			Purchase:       false,
			Payout:         false,
			Recurring:      false,
			FraudControl:   true,
			CheckoutPage:   false,
		},
	},
	{
		name:  "Тест Billing 6",
		input: "testdata/billing-6.data",
		expected: BillingData{
			CreateCustomer: false,
			Purchase:       false,
			Payout:         false,
			Recurring:      false,
			FraudControl:   false,
			CheckoutPage:   true,
		},
	},
	{
		name:  "Тест Billing All",
		input: "testdata/billing-all.data",
		expected: BillingData{
			CreateCustomer: true,
			Purchase:       true,
			Payout:         true,
			Recurring:      true,
			FraudControl:   true,
			CheckoutPage:   true,
		},
	},
	{
		name:     "Тест Billing err1",
		input:    "testdata/billing-err1.data",
		expected: BillingData{},
	},
	{
		name:     "Тест Billing err2",
		input:    "testdata/billing-err2.data",
		expected: BillingData{},
	},
}

func TestStatusBilling(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StatusBilling(test.input)
			require.Equal(t, test.expected, result)
		})
	}
}
