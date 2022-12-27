package support

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusSupport(t *testing.T) {
	t.Run("Тест Support http коды 200 и 500", func(t *testing.T) {
		handlerHttp := &TestHandler{}

		ts := httptest.NewServer(handlerHttp)
		defer ts.Close()

		result := StatusSupport(ts.URL + "/support/ok")
		require.Equal(t, data, result)

		empty := StatusSupport(ts.URL + "/support/fail")
		require.Equal(t, []SupportData{}, empty)
	})
}

var data = []SupportData{
	{Topic: "SMS", ActiveTickets: 5},
	{Topic: "MMS", ActiveTickets: 5},
	{Topic: "Email", ActiveTickets: 0},
	{Topic: "Billing", ActiveTickets: 4},
	{Topic: "Create account", ActiveTickets: 3},
	{Topic: "API", ActiveTickets: 7},
	{Topic: "Marketing", ActiveTickets: 7},
	{Topic: "Privacy", ActiveTickets: 3},
	{Topic: "GDPR", ActiveTickets: 7},
	{Topic: "Other", ActiveTickets: 7},
}

type TestHandler struct{}

func (m *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/support/ok":
		response, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	case "/support/fail":
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
