package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSsoGettokenRequest() *OapiSsoGettokenRequest {
	return &OapiSsoGettokenRequest{}
}

type OapiSsoGettokenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSsoGettokenResponse
	Corpid          string
	Corpsecret      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSsoGettokenRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiSsoGettokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSsoGettokenRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiSsoGettokenRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiSsoGettokenRequest) SetCorpsecret(corpsecret2 string) {
	this.Corpsecret = corpsecret2
}
func (this *OapiSsoGettokenRequest) GetCorpsecret() string {
	return this.Corpsecret
}
func (this *OapiSsoGettokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sso.gettoken"
}
func (this *OapiSsoGettokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSsoGettokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSsoGettokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSsoGettokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSsoGettokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corpid"] = this.Corpid
	txtParams["corpsecret"] = this.Corpsecret
	return txtParams
}
func (this *OapiSsoGettokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSsoGettokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSsoGettokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSsoGettokenResponse struct {
	taobao.TaobaoResponse
	AccessToken string `json:"access_token,omitempty"`
	Errcode     int64  `json:"errcode,omitempty"`
	Errmsg      string `json:"errmsg,omitempty"`
}
