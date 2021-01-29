package taobao

type ApiRuleError struct {
	ErrCode    string
	ErrMsg     string
	SubErrCode string
	SubErrMsg  string
}

func (are *ApiRuleError) Error() string {
	return ""
}
