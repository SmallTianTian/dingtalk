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
