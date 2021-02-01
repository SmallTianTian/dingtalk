package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserGetuserinfoRequest() *OapiUserGetuserinfoRequest {
	return &OapiUserGetuserinfoRequest{}
}

type OapiUserGetuserinfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetuserinfoResponse
	Code            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserGetuserinfoRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserGetuserinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetuserinfoRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiUserGetuserinfoRequest) GetCode() string {
	return this.Code
}
func (this *OapiUserGetuserinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.getuserinfo"
}
func (this *OapiUserGetuserinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetuserinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetuserinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetuserinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetuserinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.ERROR_CODE] = this.Code
	return txtParams
}
func (this *OapiUserGetuserinfoRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserGetuserinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetuserinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetuserinfoResponse struct {
	taobao.TaobaoResponse
	DeviceId string `json:"deviceId,omitempty"`

	IsSys    bool   `json:"is_sys,omitempty"`
	SysLevel string `json:"sys_level,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
