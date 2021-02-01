package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRetailUserTokenGenerateRequest() *OapiRetailUserTokenGenerateRequest {
	return &OapiRetailUserTokenGenerateRequest{}
}

type OapiRetailUserTokenGenerateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRetailUserTokenGenerateResponse
	Channel         string
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRetailUserTokenGenerateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailUserTokenGenerateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailUserTokenGenerateRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiRetailUserTokenGenerateRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiRetailUserTokenGenerateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRetailUserTokenGenerateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRetailUserTokenGenerateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.user.token.generate"
}
func (this *OapiRetailUserTokenGenerateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailUserTokenGenerateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailUserTokenGenerateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailUserTokenGenerateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailUserTokenGenerateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["channel"] = this.Channel
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRetailUserTokenGenerateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRetailUserTokenGenerateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailUserTokenGenerateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRetailUserTokenGenerateResponse struct {
	taobao.TaobaoResponse
	Result  UserBindDto `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
