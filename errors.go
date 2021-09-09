package netatmo

import (
	"fmt"
	"net/http"
	"strings"
)

/*
	https://dev.netatmo.com/apidocumentation/general
*/

// HTTPStatusGenericError is used to represent a non 200 HTTP error
type HTTPStatusGenericError struct {
	HTTPCode    int    ``
	NetatmoCode int    `json:"code"`
	Message     string `json:"message"`
}

func (hsge HTTPStatusGenericError) Error() string {
	tmp := fmt.Sprintf("%d %s: error code %d",
		hsge.HTTPCode, http.StatusText(hsge.HTTPCode), hsge.NetatmoCode)
	if hsge.Message != "" {
		tmp += ": " + hsge.Message
	}
	return tmp
}

// HTTPStatusOKErrors represents https://dev.netatmo.com/apidocumentation/general#status-ok
type HTTPStatusOKErrors []HTTPStatusOKError

func (hsoes HTTPStatusOKErrors) Error() string {
	var buffStr strings.Builder
	// header
	buffStr.WriteString(fmt.Sprintf("%d %s: %d error", http.StatusOK, http.StatusText(http.StatusOK), len(hsoes)))
	if len(hsoes) > 1 {
		buffStr.WriteString("s")
	}
	buffStr.WriteString(": ")
	// body
	tmp := make([]string, len(hsoes))
	for index, hsoe := range hsoes {
		tmp[index] = hsoe.Error()
	}
	buffStr.WriteString(strings.Join(tmp, ", "))
	// done
	return buffStr.String()
}

// HTTPStatusOKError represents a single API error while the HTTP request has returned 200
type HTTPStatusOKError struct {
	Code     int    `json:"code"`
	DeviceID string `json:"id"`
}

var statusOKErrors = map[int]string{
	1:  "unknown_error",
	2:  "internal_error",
	3:  "parser_error",
	5:  "command_invalid_params",
	6:  "device_unreachable",
	7:  "command_error",
	8:  "battery_level",
	14: "busy",
	19: "module_unreachable",
	23: "nothing_to_modify",
	27: "temporarily_banned",
}

func (hsoe HTTPStatusOKError) Error() string {
	return fmt.Sprintf("error code %d ('%s') for device '%s'", hsoe.Code, statusOKErrors[hsoe.Code], hsoe.DeviceID)
}

// UnexpectedHTTPCode will be used for any unexpected HTTP error codes
type UnexpectedHTTPCode struct {
	HTTPCode int
	Body     []byte
}

func (uhc UnexpectedHTTPCode) Error() string {
	return fmt.Sprintf("%d %s (body size: %d)", uhc.HTTPCode, http.StatusText(uhc.HTTPCode), len(uhc.Body))
}
