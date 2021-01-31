package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceGetCorpTokenRequest() *OapiServiceGetCorpTokenRequest {
	return &OapiServiceGetCorpTokenRequest{}
}

type OapiServiceGetCorpTokenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceGetCorpTokenResponse
	AuthCorpid      string
	PermanentCode   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceGetCorpTokenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceGetCorpTokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceGetCorpTokenRequest) SetAuthCorpid(authCorpid2 string) {
	this.AuthCorpid = authCorpid2
}
func (this *OapiServiceGetCorpTokenRequest) GetAuthCorpid() string {
	return this.AuthCorpid
}
func (this *OapiServiceGetCorpTokenRequest) SetPermanentCode(permanentCode2 string) {
	this.PermanentCode = permanentCode2
}
func (this *OapiServiceGetCorpTokenRequest) GetPermanentCode() string {
	return this.PermanentCode
}
func (this *OapiServiceGetCorpTokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.get_corp_token"
}
func (this *OapiServiceGetCorpTokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceGetCorpTokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceGetCorpTokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceGetCorpTokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceGetCorpTokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["auth_corpid"] = this.AuthCorpid
	txtParams["permanent_code"] = this.PermanentCode
	return txtParams
}
func (this *OapiServiceGetCorpTokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceGetCorpTokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceGetCorpTokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceGetCorpTokenResponse struct {
	taobao.TaobaoResponse
	AccessToken string `json:"access_token,omitempty"`
	Errcode     int64  `json:"errcode,omitempty"`
	Errmsg      string `json:"errmsg,omitempty"`
	ExpiresIn   int64  `json:"expires_in,omitempty"`
}
