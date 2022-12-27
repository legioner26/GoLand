package incident

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusIncident(t *testing.T) {
	t.Run("Тест Incident http коды 200 и 500", func(t *testing.T) {
		handlerHttp := &TestHandler{}

		ts := httptest.NewServer(handlerHttp)
		defer ts.Close()

		result := StatusIncident(ts.URL + "/incident/ok")
		require.Equal(t, data, result)

		empty := StatusIncident(ts.URL + "/incident/fail")
		require.Equal(t, []IncidentData{}, empty)
	})
}

var data = []IncidentData{
	{Topic: "SMS delivery in EU", Status: "closed"},
	{Topic: "MMS connection stability", Status: "closed"},
	{Topic: "Voice call connection purity", Status: "closed"},
	{Topic: "Checkout page is down", Status: "closed"},
	{Topic: "Support overload", Status: "active"},
	{Topic: "Buy phone number not working in US", Status: "closed"},
	{Topic: "API Slow latency", Status: "closed"},
}

type TestHandler struct{}

func (m *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/incident/ok":
		response, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	case "/incident/fail":
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
