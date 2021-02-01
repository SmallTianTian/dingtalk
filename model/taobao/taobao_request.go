package taobao

type TaobaoRequest interface {
	Check() *ApiRuleError
	GetApiMethodName() string
	GetBatchApiOrder() int
	GetBatchApiSession() string
	GetHeaderMap() map[string]string
	GetTargetAppKey() string
	GetTextParams() map[string]interface{}
	GetTimestamp() int64
	GetTopApiCallType() string
	GetTopApiFormat() string
	GetTopApiVersion() string
	GetTopContentType() string
	GetTopHttpMethod() string
	GetTopResponseType() string
	SetBatchApiOrder(i int)
	SetBatchApiSession(str string)
	SetTopApiCallType(str string)
	SetTopApiFormat(str string)
	SetTopApiVersion(str string)
	SetTopContentType(str string)
	SetTopHttpMethod(str string)
	SetTopResponseType(str string)

	// Response
	GetRespInstance() interface{}
	GetTaobaoResp() *TaobaoResponse
}

type SimpleTaobaoRequest struct{}

func (str *SimpleTaobaoRequest) Check() *ApiRuleError {
	return nil
}

func (str *SimpleTaobaoRequest) GetApiMethodName() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetBatchApiOrder() int {
	return 0
}

func (str *SimpleTaobaoRequest) GetBatchApiSession() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetHeaderMap() map[string]string {
	return nil
}

func (str *SimpleTaobaoRequest) GetTargetAppKey() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetTextParams() map[string]interface{} {
	return nil
}

func (str *SimpleTaobaoRequest) GetTimestamp() int64 {
	return 0
}

func (str *SimpleTaobaoRequest) GetTopApiCallType() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetTopApiFormat() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetTopApiVersion() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetTopContentType() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetTopHttpMethod() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetTopResponseType() string {
	return ""
}

func (str *SimpleTaobaoRequest) GetRespInstance() interface{} {
	return nil
}

func (str *SimpleTaobaoRequest) GetTaobaoResp() *TaobaoResponse {
	return nil
}

func (str *SimpleTaobaoRequest) SetBatchApiOrder(i int)      {}
func (str *SimpleTaobaoRequest) SetBatchApiSession(s string) {}
func (str *SimpleTaobaoRequest) SetTopApiCallType(s string)  {}
func (str *SimpleTaobaoRequest) SetTopApiFormat(s string)    {}
func (str *SimpleTaobaoRequest) SetTopApiVersion(s string)   {}
func (str *SimpleTaobaoRequest) SetTopContentType(s string)  {}
func (str *SimpleTaobaoRequest) SetTopHttpMethod(s string)   {}
func (str *SimpleTaobaoRequest) SetTopResponseType(s string) {}
