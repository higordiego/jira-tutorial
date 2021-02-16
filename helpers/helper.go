package helpers

import (
	"encoding/base64"
	"fmt"
	"time"
)

// BasicAuth - create basic authenticate
func BasicAuth(email, token string) string {
	auth := email + ":" + token
	return fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(auth)))
}

// NowDate - func date now parse to string
func NowDate() string {
	t := time.Now()
	return FormatDate(t)
}

// FormatDate - Parse data to string
func FormatDate(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func ConvertHour(value float64) float64 {
	return value / 3600
}
