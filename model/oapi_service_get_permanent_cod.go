package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceGetPermanentCodeRequest() *OapiServiceGetPermanentCodeRequest {
	return &OapiServiceGetPermanentCodeRequest{}
}

type OapiServiceGetPermanentCodeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceGetPermanentCodeResponse
	TmpAuthCode     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceGetPermanentCodeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceGetPermanentCodeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceGetPermanentCodeRequest) SetTmpAuthCode(tmpAuthCode2 string) {
	this.TmpAuthCode = tmpAuthCode2
}
func (this *OapiServiceGetPermanentCodeRequest) GetTmpAuthCode() string {
	return this.TmpAuthCode
}
func (this *OapiServiceGetPermanentCodeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.get_permanent_code"
}
func (this *OapiServiceGetPermanentCodeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceGetPermanentCodeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceGetPermanentCodeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceGetPermanentCodeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceGetPermanentCodeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["tmp_auth_code"] = this.TmpAuthCode
	return txtParams
}
func (this *OapiServiceGetPermanentCodeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceGetPermanentCodeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceGetPermanentCodeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceGetPermanentCodeResponse struct {
	taobao.TaobaoResponse
	AuthCorpInfo    AuthCorpInfo `json:"auth_corp_info,omitempty"`
	ChPermanentCode string       `json:"ch_permanent_code,omitempty"`

	PermanentCode string `json:"permanent_code,omitempty"`
}
