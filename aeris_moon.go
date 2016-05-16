package suriya

import (
	"time"
)

var phaseCodes = []string{"new", "waxing", "full", "waning"}

type AerisMoon struct {
	Timestamp   int       `json:"timestamp"`
	DateTimeISO time.Time `json:"dateTimeISO"`
	Code        int       `json:"code"`
	Name        string    `json:"name"`
}

type AerisError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type AerisResp struct {
	Success  bool       `json:"success"`
	Error    AerisError `json:"error"`
	Response []AerisMoon
}
