package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiGettokenRequest() *OapiGettokenRequest {
	return &OapiGettokenRequest{}
}

type OapiGettokenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiGettokenResponse
	Appkey          string
	Appsecret       string
	Corpid          string
	Corpsecret      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiGettokenRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiGettokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiGettokenRequest) SetAppkey(appkey2 string) {
	this.Appkey = appkey2
}
func (this *OapiGettokenRequest) GetAppkey() string {
	return this.Appkey
}
func (this *OapiGettokenRequest) SetAppsecret(appsecret2 string) {
	this.Appsecret = appsecret2
}
func (this *OapiGettokenRequest) GetAppsecret() string {
	return this.Appsecret
}
func (this *OapiGettokenRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiGettokenRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiGettokenRequest) SetCorpsecret(corpsecret2 string) {
	this.Corpsecret = corpsecret2
}
func (this *OapiGettokenRequest) GetCorpsecret() string {
	return this.Corpsecret
}
func (this *OapiGettokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.gettoken"
}
func (this *OapiGettokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiGettokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiGettokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiGettokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiGettokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["appkey"] = this.Appkey
	txtParams["appsecret"] = this.Appsecret
	txtParams["corpid"] = this.Corpid
	txtParams["corpsecret"] = this.Corpsecret
	return txtParams
}
func (this *OapiGettokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiGettokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiGettokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiGettokenResponse struct {
	taobao.TaobaoResponse
	AccessToken string `json:"access_token,omitempty"`
	Errcode     int64  `json:"errcode,omitempty"`
	Errmsg      string `json:"errmsg,omitempty"`
	ExpiresIn   int64  `json:"expires_in,omitempty"`
}
