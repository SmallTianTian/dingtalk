package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRetailUserTokenCheckRequest() *OapiRetailUserTokenCheckRequest {
	return &OapiRetailUserTokenCheckRequest{}
}

type OapiRetailUserTokenCheckRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRetailUserTokenCheckResponse
	Channel         string
	Token           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRetailUserTokenCheckRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailUserTokenCheckRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailUserTokenCheckRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiRetailUserTokenCheckRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiRetailUserTokenCheckRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiRetailUserTokenCheckRequest) GetToken() string {
	return this.Token
}
func (this *OapiRetailUserTokenCheckRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.user.token.check"
}
func (this *OapiRetailUserTokenCheckRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailUserTokenCheckRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailUserTokenCheckRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailUserTokenCheckRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailUserTokenCheckRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["channel"] = this.Channel
	txtParams["token"] = this.Token
	return txtParams
}
func (this *OapiRetailUserTokenCheckRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRetailUserTokenCheckRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailUserTokenCheckRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRetailUserTokenCheckResponse struct {
	taobao.TaobaoResponse
	Result  UserBindDto `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
