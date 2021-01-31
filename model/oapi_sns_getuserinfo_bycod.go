package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSnsGetuserinfoBycodeRequest() *OapiSnsGetuserinfoBycodeRequest {
	return &OapiSnsGetuserinfoBycodeRequest{}
}

type OapiSnsGetuserinfoBycodeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSnsGetuserinfoBycodeResponse
	TmpAuthCode     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSnsGetuserinfoBycodeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSnsGetuserinfoBycodeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsGetuserinfoBycodeRequest) SetTmpAuthCode(tmpAuthCode2 string) {
	this.TmpAuthCode = tmpAuthCode2
}
func (this *OapiSnsGetuserinfoBycodeRequest) GetTmpAuthCode() string {
	return this.TmpAuthCode
}
func (this *OapiSnsGetuserinfoBycodeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.getuserinfo_bycode"
}
func (this *OapiSnsGetuserinfoBycodeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsGetuserinfoBycodeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsGetuserinfoBycodeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsGetuserinfoBycodeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsGetuserinfoBycodeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["tmp_auth_code"] = this.TmpAuthCode
	return txtParams
}
func (this *OapiSnsGetuserinfoBycodeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSnsGetuserinfoBycodeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsGetuserinfoBycodeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsGetuserinfoBycodeResponse struct {
	taobao.TaobaoResponse
	Errcode  int64    `json:"errcode,omitempty"`
	Errmsg   string   `json:"errmsg,omitempty"`
	UserInfo UserInfo `json:"user_info,omitempty"`
}
type UserInfo struct {
	MainOrgAuthHighLevel bool   `json:"main_org_auth_high_level,omitempty"`
	Nick                 string `json:"nick,omitempty"`
	Openid               string `json:"openid,omitempty"`
	Unionid              string `json:"unionid,omitempty"`
}
