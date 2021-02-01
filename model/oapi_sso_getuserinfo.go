package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSsoGetuserinfoRequest() *OapiSsoGetuserinfoRequest {
	return &OapiSsoGetuserinfoRequest{}
}

type OapiSsoGetuserinfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSsoGetuserinfoResponse
	AccessToken     string
	Code            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSsoGetuserinfoRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiSsoGetuserinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSsoGetuserinfoRequest) SetAccessToken(accessToken2 string) {
	this.AccessToken = accessToken2
}
func (this *OapiSsoGetuserinfoRequest) GetAccessToken() string {
	return this.AccessToken
}
func (this *OapiSsoGetuserinfoRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiSsoGetuserinfoRequest) GetCode() string {
	return this.Code
}
func (this *OapiSsoGetuserinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sso.getuserinfo"
}
func (this *OapiSsoGetuserinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSsoGetuserinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSsoGetuserinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSsoGetuserinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSsoGetuserinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[ACCESS_TOKEN] = this.AccessToken
	txtParams[taobao.ERROR_CODE] = this.Code
	return txtParams
}
func (this *OapiSsoGetuserinfoRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSsoGetuserinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSsoGetuserinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSsoGetuserinfoResponse struct {
	taobao.TaobaoResponse
	CorpInfo CorpInfo `json:"corp_info,omitempty"`

	IsSys    bool     `json:"is_sys,omitempty"`
	UserInfo UserInfo `json:"user_info,omitempty"`
}
type CorpInfo struct {
	CorpName string `json:"corp_name,omitempty"`
	Corpid   string `json:"corpid,omitempty"`
}
