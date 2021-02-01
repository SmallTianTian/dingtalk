package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCallBackGetCallBackRequest() *OapiCallBackGetCallBackRequest {
	return &OapiCallBackGetCallBackRequest{}
}

type OapiCallBackGetCallBackRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallBackGetCallBackResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallBackGetCallBackRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiCallBackGetCallBackRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallBackGetCallBackRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call_back.get_call_back"
}
func (this *OapiCallBackGetCallBackRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallBackGetCallBackRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallBackGetCallBackRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallBackGetCallBackRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallBackGetCallBackRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCallBackGetCallBackRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCallBackGetCallBackRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallBackGetCallBackRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallBackGetCallBackResponse struct {
	taobao.TaobaoResponse
	AesKey      string   `json:"aes_key,omitempty"`
	CallBackTag []string `json:"call_back_tag,omitempty"`

	Token string `json:"token,omitempty"`
	Url   string `json:"url,omitempty"`
}
