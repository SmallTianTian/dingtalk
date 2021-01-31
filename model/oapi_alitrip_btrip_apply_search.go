package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripApplySearchRequest() *OapiAlitripBtripApplySearchRequest {
	return &OapiAlitripBtripApplySearchRequest{}
}

type OapiAlitripBtripApplySearchRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripApplySearchResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripApplySearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripApplySearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripApplySearchRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripApplySearchRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripApplySearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.apply.search"
}
func (this *OapiAlitripBtripApplySearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripApplySearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripApplySearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripApplySearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripApplySearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripApplySearchRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripApplySearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripApplySearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripApplySearchResponse struct {
	taobao.TaobaoResponse
	Module  []OpenApplyRs `json:"module,omitempty"`
	Success bool          `json:"success,omitempty"`
}
