package netatmo

import (
	"fmt"
	"net/http"
	"strings"
)

/*
	https://dev.netatmo.com/apidocumentation/general
*/

type HTTPStatusGenericError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (hsge HTTPStatusGenericError) generateError(httpcode int) string {
	tmp := fmt.Sprintf("%d %s: error code %d",
		httpcode, http.StatusText(httpcode), hsge.Code)
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

// HTTPStatusBadRequestError represents https://dev.netatmo.com/apidocumentation/general#status-bad-request
type HTTPStatusBadRequestError HTTPStatusGenericError

func (hsbre HTTPStatusBadRequestError) Error() string {
	return (HTTPStatusGenericError)(hsbre).generateError(http.StatusBadRequest)
}

// HTTPStatusUnauthorizedError represents https://dev.netatmo.com/apidocumentation/general#status-unauthorized
type HTTPStatusUnauthorizedError HTTPStatusGenericError

func (hsue HTTPStatusUnauthorizedError) Error() string {
	return (HTTPStatusGenericError)(hsue).generateError(http.StatusUnauthorized)
}

// HTTPStatusForbiddenError represents https://dev.netatmo.com/apidocumentation/general#status-forbidden
type HTTPStatusForbiddenError HTTPStatusGenericError

func (hsfe HTTPStatusForbiddenError) Error() string {
	return (HTTPStatusGenericError)(hsfe).generateError(http.StatusForbidden)
}

// HTTPStatusNotFoundError represents https://dev.netatmo.com/apidocumentation/general#status-not-found
type HTTPStatusNotFoundError HTTPStatusGenericError

func (hsnfe HTTPStatusNotFoundError) Error() string {
	return (HTTPStatusGenericError)(hsnfe).generateError(http.StatusNotFound)
}

// HTTPStatusNotAcceptableError represents https://dev.netatmo.com/apidocumentation/general#status-not-acceptable
type HTTPStatusNotAcceptableError HTTPStatusGenericError

func (hsnae HTTPStatusNotAcceptableError) Error() string {
	return (HTTPStatusGenericError)(hsnae).generateError(http.StatusNotAcceptable)
}

// HTTPStatusInternalServerErrorError represents https://dev.netatmo.com/apidocumentation/general#internal-error
type HTTPStatusInternalServerErrorError HTTPStatusGenericError

func (hsisee HTTPStatusInternalServerErrorError) Error() string {
	return (HTTPStatusGenericError)(hsisee).generateError(http.StatusInternalServerError)
}

// UnexpectedHTTPCode will be used for any unexpected HTTP error codes
type UnexpectedHTTPCode struct {
	Code int
	Body []byte
}

func (uhc UnexpectedHTTPCode) Error() string {
	return fmt.Sprintf("%d %s (body size: %d)", uhc.Code, http.StatusText(uhc.Code), len(uhc.Body))
}
