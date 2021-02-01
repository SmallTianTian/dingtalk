package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCallBackDeleteCallBackRequest() *OapiCallBackDeleteCallBackRequest {
	return &OapiCallBackDeleteCallBackRequest{}
}

type OapiCallBackDeleteCallBackRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallBackDeleteCallBackResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallBackDeleteCallBackRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiCallBackDeleteCallBackRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallBackDeleteCallBackRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call_back.delete_call_back"
}
func (this *OapiCallBackDeleteCallBackRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallBackDeleteCallBackRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallBackDeleteCallBackRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallBackDeleteCallBackRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallBackDeleteCallBackRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCallBackDeleteCallBackRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCallBackDeleteCallBackRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallBackDeleteCallBackRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallBackDeleteCallBackResponse struct {
	taobao.TaobaoResponse
}
