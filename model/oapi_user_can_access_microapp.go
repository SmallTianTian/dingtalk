package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserCanAccessMicroappRequest() *OapiUserCanAccessMicroappRequest {
	return &OapiUserCanAccessMicroappRequest{}
}

type OapiUserCanAccessMicroappRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserCanAccessMicroappResponse
	AppId           int64
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *OapiUserCanAccessMicroappRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserCanAccessMicroappRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserCanAccessMicroappRequest) SetAppId(appId2 int64) {
	this.AppId = appId2
}
func (this *OapiUserCanAccessMicroappRequest) GetAppId() int64 {
	return this.AppId
}
func (this *OapiUserCanAccessMicroappRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiUserCanAccessMicroappRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiUserCanAccessMicroappRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.can_access_microapp"
}
func (this *OapiUserCanAccessMicroappRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserCanAccessMicroappRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserCanAccessMicroappRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserCanAccessMicroappRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserCanAccessMicroappRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["appId"] = this.AppId
	txtParams["userId"] = this.UserId
	return txtParams
}
func (this *OapiUserCanAccessMicroappRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserCanAccessMicroappRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserCanAccessMicroappRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserCanAccessMicroappResponse struct {
	taobao.TaobaoResponse
	CanAccess bool   `json:"canAccess,omitempty"`
	Errcode   int64  `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
}
