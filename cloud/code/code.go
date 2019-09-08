package code

import (
	"net/http"
)

type Code string

const (
	InternalServerError Code = "INTERNAL_SERVER_ERROR"
	NotFound                 = "NOT_FOUND"
	MethodNotAllowed         = "METHOD_NOT_ALLOWED"

	InvalidData = "INVALID_DATA"

	UserExist       = "USER_EXIST"
	UserNotExist    = "USER_NOT_EXIST"
	InvalidPassword = "INVALID_PASSWORD"
	InvalidToken    = "INVALID_TOKEN"

	RuleNotExist   = "RULE_NOT_EXIST"
	RecordNotExit  = "RECORD_NOT_EXIST"
	DeviceNotExist = "DEVICE_NOT_EXIST"

	BlackListNotExist  = "BLACK_LIST_NOT_EXIST"
	VeinRecordNotExist = "VEIN_RECORD_NOT_EXIST"

	IdentityDataIvalid = "IDENTITY_DATA_INVALID"
	IdentityNotExist   = "IDENTITY_NOT_EXIST"
	IdentityExist      = "IDENTITY_EXIST"
)

var codeStatusMap = map[Code]int{
	InternalServerError: http.StatusInternalServerError,
	NotFound:            http.StatusNotFound,
	MethodNotAllowed:    http.StatusMethodNotAllowed,
	InvalidData:         http.StatusBadRequest,
	UserExist:           http.StatusBadRequest,
	UserNotExist:        http.StatusNotFound,
	InvalidPassword:     http.StatusBadRequest,
}
