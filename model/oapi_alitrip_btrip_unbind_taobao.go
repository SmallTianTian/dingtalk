package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripUnbindTaobaoRequest() *OapiAlitripBtripUnbindTaobaoRequest {
	return &OapiAlitripBtripUnbindTaobaoRequest{}
}

type OapiAlitripBtripUnbindTaobaoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripUnbindTaobaoResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripUnbindTaobaoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.unbind.taobao"
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripUnbindTaobaoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripUnbindTaobaoResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
