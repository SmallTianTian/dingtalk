package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavTelconfGetRequest() *OapiKacDatavTelconfGetRequest {
	return &OapiKacDatavTelconfGetRequest{}
}

type OapiKacDatavTelconfGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavTelconfGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavTelconfGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavTelconfGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavTelconfGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavTelconfGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavTelconfGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.telconf.get"
}
func (this *OapiKacDatavTelconfGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavTelconfGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavTelconfGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavTelconfGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavTelconfGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavTelconfGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavTelconfGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavTelconfGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacDatavTelconfGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                        `json:"errcode,omitempty"`
	Errmsg  string                       `json:"errmsg,omitempty"`
	Result  TelConferenceSummaryResponse `json:"result,omitempty"`
}
type TelConferenceSummaryResponse struct {
	CallCount       int64  `json:"call_count,omitempty"`
	CallDuration    string `json:"call_duration,omitempty"`
	CallDurationMin string `json:"call_duration_min,omitempty"`
	CallJoinPv      int64  `json:"call_join_pv,omitempty"`
}
