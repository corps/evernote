// This file is automatically generated. Do not modify.

package edamerrors

import (
	"fmt"
	"strconv"
)

type EDAMErrorCode int32

const (
	EDAMErrorCodeAuthExpired          EDAMErrorCode = 9
	EDAMErrorCodeBadDataFormat        EDAMErrorCode = 2
	EDAMErrorCodeDataConflict         EDAMErrorCode = 10
	EDAMErrorCodeDataRequired         EDAMErrorCode = 5
	EDAMErrorCodeEnmlValidation       EDAMErrorCode = 11
	EDAMErrorCodeInternalError        EDAMErrorCode = 4
	EDAMErrorCodeInvalidAuth          EDAMErrorCode = 8
	EDAMErrorCodeLenTooLong           EDAMErrorCode = 14
	EDAMErrorCodeLenTooShort          EDAMErrorCode = 13
	EDAMErrorCodeLimitReached         EDAMErrorCode = 6
	EDAMErrorCodePermissionDenied     EDAMErrorCode = 3
	EDAMErrorCodeQuotaReached         EDAMErrorCode = 7
	EDAMErrorCodeRateLimitReached     EDAMErrorCode = 19
	EDAMErrorCodeShardUnavailable     EDAMErrorCode = 12
	EDAMErrorCodeTakenDown            EDAMErrorCode = 18
	EDAMErrorCodeTooFew               EDAMErrorCode = 15
	EDAMErrorCodeTooMany              EDAMErrorCode = 16
	EDAMErrorCodeUnknown              EDAMErrorCode = 1
	EDAMErrorCodeUnsupportedOperation EDAMErrorCode = 17
)

var (
	EDAMErrorCodeByName = map[string]EDAMErrorCode{
		"EDAMErrorCode.AUTH_EXPIRED":          EDAMErrorCodeAuthExpired,
		"EDAMErrorCode.BAD_DATA_FORMAT":       EDAMErrorCodeBadDataFormat,
		"EDAMErrorCode.DATA_CONFLICT":         EDAMErrorCodeDataConflict,
		"EDAMErrorCode.DATA_REQUIRED":         EDAMErrorCodeDataRequired,
		"EDAMErrorCode.ENML_VALIDATION":       EDAMErrorCodeEnmlValidation,
		"EDAMErrorCode.INTERNAL_ERROR":        EDAMErrorCodeInternalError,
		"EDAMErrorCode.INVALID_AUTH":          EDAMErrorCodeInvalidAuth,
		"EDAMErrorCode.LEN_TOO_LONG":          EDAMErrorCodeLenTooLong,
		"EDAMErrorCode.LEN_TOO_SHORT":         EDAMErrorCodeLenTooShort,
		"EDAMErrorCode.LIMIT_REACHED":         EDAMErrorCodeLimitReached,
		"EDAMErrorCode.PERMISSION_DENIED":     EDAMErrorCodePermissionDenied,
		"EDAMErrorCode.QUOTA_REACHED":         EDAMErrorCodeQuotaReached,
		"EDAMErrorCode.RATE_LIMIT_REACHED":    EDAMErrorCodeRateLimitReached,
		"EDAMErrorCode.SHARD_UNAVAILABLE":     EDAMErrorCodeShardUnavailable,
		"EDAMErrorCode.TAKEN_DOWN":            EDAMErrorCodeTakenDown,
		"EDAMErrorCode.TOO_FEW":               EDAMErrorCodeTooFew,
		"EDAMErrorCode.TOO_MANY":              EDAMErrorCodeTooMany,
		"EDAMErrorCode.UNKNOWN":               EDAMErrorCodeUnknown,
		"EDAMErrorCode.UNSUPPORTED_OPERATION": EDAMErrorCodeUnsupportedOperation,
	}
	EDAMErrorCodeByValue = map[EDAMErrorCode]string{
		EDAMErrorCodeAuthExpired:          "EDAMErrorCode.AUTH_EXPIRED",
		EDAMErrorCodeBadDataFormat:        "EDAMErrorCode.BAD_DATA_FORMAT",
		EDAMErrorCodeDataConflict:         "EDAMErrorCode.DATA_CONFLICT",
		EDAMErrorCodeDataRequired:         "EDAMErrorCode.DATA_REQUIRED",
		EDAMErrorCodeEnmlValidation:       "EDAMErrorCode.ENML_VALIDATION",
		EDAMErrorCodeInternalError:        "EDAMErrorCode.INTERNAL_ERROR",
		EDAMErrorCodeInvalidAuth:          "EDAMErrorCode.INVALID_AUTH",
		EDAMErrorCodeLenTooLong:           "EDAMErrorCode.LEN_TOO_LONG",
		EDAMErrorCodeLenTooShort:          "EDAMErrorCode.LEN_TOO_SHORT",
		EDAMErrorCodeLimitReached:         "EDAMErrorCode.LIMIT_REACHED",
		EDAMErrorCodePermissionDenied:     "EDAMErrorCode.PERMISSION_DENIED",
		EDAMErrorCodeQuotaReached:         "EDAMErrorCode.QUOTA_REACHED",
		EDAMErrorCodeRateLimitReached:     "EDAMErrorCode.RATE_LIMIT_REACHED",
		EDAMErrorCodeShardUnavailable:     "EDAMErrorCode.SHARD_UNAVAILABLE",
		EDAMErrorCodeTakenDown:            "EDAMErrorCode.TAKEN_DOWN",
		EDAMErrorCodeTooFew:               "EDAMErrorCode.TOO_FEW",
		EDAMErrorCodeTooMany:              "EDAMErrorCode.TOO_MANY",
		EDAMErrorCodeUnknown:              "EDAMErrorCode.UNKNOWN",
		EDAMErrorCodeUnsupportedOperation: "EDAMErrorCode.UNSUPPORTED_OPERATION",
	}
)

func (e EDAMErrorCode) String() string {
	name := EDAMErrorCodeByValue[e]
	if name == "" {
		name = fmt.Sprintf("Unknown enum value EDAMErrorCode(%d)", e)
	}
	return name
}

func (e EDAMErrorCode) MarshalJSON() ([]byte, error) {
	name := EDAMErrorCodeByValue[e]
	if name == "" {
		name = strconv.Itoa(int(e))
	}
	return []byte("\"" + name + "\""), nil
}

func (e *EDAMErrorCode) UnmarshalJSON(b []byte) error {
	st := string(b)
	if st[0] == '"' {
		*e = EDAMErrorCode(EDAMErrorCodeByName[st[1:len(st)-1]])
		return nil
	}
	i, err := strconv.Atoi(st)
	*e = EDAMErrorCode(i)
	return err
}

type EDAMNotFoundException struct {
	Identifier *string `thrift:"1" json:"identifier,omitempty"`
	Key        *string `thrift:"2" json:"key,omitempty"`
}

func (e *EDAMNotFoundException) Error() string {
	return fmt.Sprintf("EDAMNotFoundException{Identifier: %+v, Key: %+v}", e.Identifier, e.Key)
}

type EDAMSystemException struct {
	ErrorCode         *EDAMErrorCode `thrift:"1,required" json:"errorCode"`
	Message           *string        `thrift:"2" json:"message,omitempty"`
	RateLimitDuration *int32         `thrift:"3" json:"rateLimitDuration,omitempty"`
}

func (e *EDAMSystemException) Error() string {
	return fmt.Sprintf("EDAMSystemException{ErrorCode: %+v, Message: %+v, RateLimitDuration: %+v}", e.ErrorCode, e.Message, e.RateLimitDuration)
}

type EDAMUserException struct {
	ErrorCode *EDAMErrorCode `thrift:"1,required" json:"errorCode"`
	Parameter *string        `thrift:"2" json:"parameter,omitempty"`
}

func (e *EDAMUserException) Error() string {
	return fmt.Sprintf("EDAMUserException{ErrorCode: %+v, Parameter: %+v}", e.ErrorCode, e.Parameter)
}
