package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceGetSuiteTokenRequest() *OapiServiceGetSuiteTokenRequest {
	return &OapiServiceGetSuiteTokenRequest{}
}

type OapiServiceGetSuiteTokenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceGetSuiteTokenResponse
	SuiteKey        string
	SuiteSecret     string
	SuiteTicket     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceGetSuiteTokenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceGetSuiteTokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceGetSuiteTokenRequest) SetSuiteKey(suiteKey2 string) {
	this.SuiteKey = suiteKey2
}
func (this *OapiServiceGetSuiteTokenRequest) GetSuiteKey() string {
	return this.SuiteKey
}
func (this *OapiServiceGetSuiteTokenRequest) SetSuiteSecret(suiteSecret2 string) {
	this.SuiteSecret = suiteSecret2
}
func (this *OapiServiceGetSuiteTokenRequest) GetSuiteSecret() string {
	return this.SuiteSecret
}
func (this *OapiServiceGetSuiteTokenRequest) SetSuiteTicket(suiteTicket2 string) {
	this.SuiteTicket = suiteTicket2
}
func (this *OapiServiceGetSuiteTokenRequest) GetSuiteTicket() string {
	return this.SuiteTicket
}
func (this *OapiServiceGetSuiteTokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.get_suite_token"
}
func (this *OapiServiceGetSuiteTokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceGetSuiteTokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceGetSuiteTokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceGetSuiteTokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceGetSuiteTokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["suite_key"] = this.SuiteKey
	txtParams["suite_secret"] = this.SuiteSecret
	txtParams["suite_ticket"] = this.SuiteTicket
	return txtParams
}
func (this *OapiServiceGetSuiteTokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceGetSuiteTokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceGetSuiteTokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceGetSuiteTokenResponse struct {
	taobao.TaobaoResponse
	Errcode          int64  `json:"errcode,omitempty"`
	Errmsg           string `json:"errmsg,omitempty"`
	ExpiresIn        int64  `json:"expires_in,omitempty"`
	SuiteAccessToken string `json:"suite_access_token,omitempty"`
}
