package check

import (
	"strconv"
)

func IsCountry(code string) bool {
	countryMutex.Lock()
	_, ok := countries[code]
	countryMutex.Unlock()
	return ok
}

func GetCountryForCode(code string) string {
	countryMutex.Lock()
	country, ok := countries[code]
	countryMutex.Unlock()
	if ok {
		return country
	}
	return ""
}

func IsBandwidth(bandwidth string) bool {
	i, err := strconv.Atoi(bandwidth)
	if err != nil {
		return false
	}

	if i >= 0 && i <= 100 {
		return true
	}

	return false
}

func IsResponseTime(time string) bool {
	i, err := strconv.Atoi(time)
	if err != nil {
		return false
	}

	if i >= 0 {
		return true
	}

	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsProviderSMSandMMS(provider string) bool {
	providers := []string{"Topolo", "Rond", "Kildy"}
	return contains(providers, provider)
}

func IsProviderVoiceCall(provider string) bool {
	providers := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	return contains(providers, provider)
}

func IsProviderEmail(provider string) bool {
	providers := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL",
		"Live", "RediffMail", "GMX", "Protonmail", "Yandex", "Mail.ru"}
	return contains(providers, provider)
}

func IsPositiveInt(value string) bool {
	i, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if i >= 0 {
		return true
	}

	return false
}

func IsPositiveFloat(value string) bool {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return false
	}

	if f >= 0 {
		return true
	}

	return false
}
