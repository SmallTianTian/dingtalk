package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSnsGetSnsTokenRequest() *OapiSnsGetSnsTokenRequest {
	return &OapiSnsGetSnsTokenRequest{}
}

type OapiSnsGetSnsTokenRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSnsGetSnsTokenResponse
	Openid          string
	PersistentCode  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSnsGetSnsTokenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSnsGetSnsTokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsGetSnsTokenRequest) SetOpenid(openid2 string) {
	this.Openid = openid2
}
func (this *OapiSnsGetSnsTokenRequest) GetOpenid() string {
	return this.Openid
}
func (this *OapiSnsGetSnsTokenRequest) SetPersistentCode(persistentCode2 string) {
	this.PersistentCode = persistentCode2
}
func (this *OapiSnsGetSnsTokenRequest) GetPersistentCode() string {
	return this.PersistentCode
}
func (this *OapiSnsGetSnsTokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.get_sns_token"
}
func (this *OapiSnsGetSnsTokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsGetSnsTokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsGetSnsTokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsGetSnsTokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsGetSnsTokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["openid"] = this.Openid
	txtParams["persistent_code"] = this.PersistentCode
	return txtParams
}
func (this *OapiSnsGetSnsTokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSnsGetSnsTokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsGetSnsTokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsGetSnsTokenResponse struct {
	taobao.TaobaoResponse

	ExpiresIn int64  `json:"expires_in,omitempty"`
	SnsToken  string `json:"sns_token,omitempty"`
}
