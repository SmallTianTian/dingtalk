package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

const (
	ERROR_CODE_ARGUMENTS_INVALID = "41"
	ERROR_CODE_ARGUMENTS_MISSING = "40"
)

func CheckNotEmpty(value interface{}, fieldName string) *taobao.ApiRuleError {
	if value == nil {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_MISSING,
			ErrMsg:  fmt.Sprintf("client-error:Missing required arguments:%s.", fieldName)}
	}
	if strV, ok := value.(string); ok && strings.TrimSpace(strV) == "" {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_MISSING,
			ErrMsg:  fmt.Sprintf("client-error:Missing required arguments:%s.", fieldName)}
	}
	return nil
}

func CheckMaxLength(value string, maxLength int, fieldName string) *taobao.ApiRuleError {
	if len(value) > maxLength {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_INVALID,
			ErrMsg:  fmt.Sprintf("client-error:Invalid arguments:the string length of %s can not be larger than %d.", fieldName, maxLength),
		}
	}
	return nil
}

func CheckFileMaxLength(fileItem *os.File, maxLength int, fieldName string) *taobao.ApiRuleError {
	info, err := fileItem.Stat()
	if err != nil {
		// no file
		return nil
	}
	if fileItem != nil && info.Size() > int64(maxLength) {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_INVALID,
			ErrMsg:  fmt.Sprintf("client-error:Invalid arguments:the file size of %s can not be larger than %d.", fieldName, maxLength),
		}
	}
	return nil
}

func CheckStringMaxListSize(value string, maxSize int, fieldName string) *taobao.ApiRuleError {
	return CheckMaxListSize(strings.Split(value, ","), maxSize, fieldName)
}

func CheckMaxListSize(list []string, maxSize int, fieldName string) *taobao.ApiRuleError {
	if len(list) > maxSize {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_INVALID,
			ErrMsg:  fmt.Sprintf("client-error:Invalid arguments:the array size of %s must be less than %d.", fieldName, maxSize),
		}
	}
	return nil
}

func CheckObjectMaxListSize(value string, maxSize int, fieldName string) *taobao.ApiRuleError {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	// maybe not json
	if !strings.HasPrefix(value, "[") && !strings.HasPrefix(value, "{") {
		return nil
	}
	var ins interface{}
	if err := json.Unmarshal([]byte(value), &ins); err != nil {
		// not json
		return nil
	}
	if listIns, ok := ins.([]interface{}); ok && len(listIns) > maxSize {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_INVALID,
			ErrMsg:  fmt.Sprintf("client-error:Invalid arguments:the array size of %s must be less than %d.", fieldName, maxSize),
		}
	}
	return nil
}

func CheckMaxValue(value int, maxValue int, fieldName string) *taobao.ApiRuleError {
	if value > maxValue {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_INVALID,
			ErrMsg:  fmt.Sprintf("client-error:Invalid arguments:the value of %s can not be larger than %d.", fieldName, maxValue),
		}
	}
	return nil
}

func CheckMinValue(value, minValue int, fieldName string) *taobao.ApiRuleError {
	if value < minValue {
		return &taobao.ApiRuleError{
			ErrCode: ERROR_CODE_ARGUMENTS_INVALID,
			ErrMsg:  fmt.Sprintf("client-error:Invalid arguments:the value of %s can not be less than %d.", fieldName, minValue),
		}
	}
	return nil
}
