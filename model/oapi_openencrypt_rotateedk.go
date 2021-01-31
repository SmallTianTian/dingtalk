package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOpenencryptRotateedkRequest() *OapiOpenencryptRotateedkRequest {
	return &OapiOpenencryptRotateedkRequest{}
}

type OapiOpenencryptRotateedkRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiOpenencryptRotateedkResponse
	TopEdkRotateManual string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiOpenencryptRotateedkRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOpenencryptRotateedkRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOpenencryptRotateedkRequest) SetTopEdkRotateManual(topEdkRotateManual2 string) {
	this.TopEdkRotateManual = topEdkRotateManual2
}
func (this *OapiOpenencryptRotateedkRequest) GetTopEdkRotateManual() string {
	return this.TopEdkRotateManual
}
func (this *OapiOpenencryptRotateedkRequest) GetApiMethodName() string {
	return "dingtalk.oapi.openencrypt.rotateedk"
}
func (this *OapiOpenencryptRotateedkRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOpenencryptRotateedkRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOpenencryptRotateedkRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOpenencryptRotateedkRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOpenencryptRotateedkRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_edk_rotate_manual"] = this.TopEdkRotateManual
	return txtParams
}
func (this *OapiOpenencryptRotateedkRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOpenencryptRotateedkRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOpenencryptRotateedkRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopEdkRotateManual struct {
	Agentid   int64  `json:"agentid,omitempty"`
	Requestid string `json:"requestid,omitempty"`
	Resource  string `json:"resource,omitempty"`
}
type OapiOpenencryptRotateedkResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
