package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSnsGetuserinfoRequest() *OapiSnsGetuserinfoRequest {
	return &OapiSnsGetuserinfoRequest{}
}

type OapiSnsGetuserinfoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSnsGetuserinfoResponse
	SnsToken        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSnsGetuserinfoRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiSnsGetuserinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsGetuserinfoRequest) SetSnsToken(snsToken2 string) {
	this.SnsToken = snsToken2
}
func (this *OapiSnsGetuserinfoRequest) GetSnsToken() string {
	return this.SnsToken
}
func (this *OapiSnsGetuserinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.getuserinfo"
}
func (this *OapiSnsGetuserinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsGetuserinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsGetuserinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsGetuserinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsGetuserinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["sns_token"] = this.SnsToken
	return txtParams
}
func (this *OapiSnsGetuserinfoRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSnsGetuserinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsGetuserinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsGetuserinfoResponse struct {
	taobao.TaobaoResponse
	Errcode  int64    `json:"errcode,omitempty"`
	Errmsg   string   `json:"errmsg,omitempty"`
	UserInfo UserInfo `json:"user_info,omitempty"`
}
