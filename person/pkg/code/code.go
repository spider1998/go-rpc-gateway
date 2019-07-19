package code

import (
	"git.sdkeji.top/share/sdlib/api"
	"sdkeji/person/pkg/app"
)

const (
	UserNotExist         api.Code = "USER_NOT_EXIST"
	UserNotBindWechat    api.Code = "USER_NOT_BIND_WECHAT"
	PersonNotExist       api.Code = "PERSON_NOT_EXIST"
	InvalidPersonToken   api.Code = "INVALID_PERSON_TOKEN"
	InvalidInternalToken api.Code = "INVALID_INTERNALTOKEN "
	RuleNotExist         api.Code = "RULE_NOT_EXIST"
	DeviceKeyNotExist    api.Code = "DEVICE_KEY_NOT_EXIST"
	FormIDNotExist       api.Code = "FORM_ID_NOT_EXIST"
)

func Error(code api.Code) api.APIError {
	return api.Error(api.ModuleError{
		Module: app.System,
		Error:  code,
	})
}
