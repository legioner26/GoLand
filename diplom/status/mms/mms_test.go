package mms

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusMMS(t *testing.T) {
	t.Run("Тест MMS http коды 200 и 500", func(t *testing.T) {
		handlerHttp := &TestHandler{}

		ts := httptest.NewServer(handlerHttp)
		defer ts.Close()

		result := StatusMMS(ts.URL + "/mms/ok")
		require.Equal(t, expected, result)

		empty := StatusMMS(ts.URL + "/mms/fail")
		require.Equal(t, []MMSData{}, empty)
	})
}

var testData = []MMSData{
	{Country: "RU", Bandwidth: "98", ResponseTime: "1576", Provider: "Topolo"},
	{Country: "US", Bandwidth: "97", ResponseTime: "108", Provider: "Rond"},
	{Country: "BL", Bandwidth: "16", ResponseTime: "1085", Provider: "Kildy"},
	{Country: "ERR", Bandwidth: "18", ResponseTime: "1990", Provider: "Topolo"},
	{Country: "BL", Bandwidth: "116", ResponseTime: "1085", Provider: "Kildy"},
	{Country: "AT", Bandwidth: "17", ResponseTime: "-124", Provider: "Topolo"},
	{Country: "BG", Bandwidth: "6", ResponseTime: "1461", Provider: "Error"},
}

var expected = []MMSData{
	{Country: "RU", Bandwidth: "98", ResponseTime: "1576", Provider: "Topolo"},
	{Country: "US", Bandwidth: "97", ResponseTime: "108", Provider: "Rond"},
	{Country: "BL", Bandwidth: "16", ResponseTime: "1085", Provider: "Kildy"},
}

type TestHandler struct{}

func (m *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/mms/ok":
		response, _ := json.Marshal(testData)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	case "/mms/fail":
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
