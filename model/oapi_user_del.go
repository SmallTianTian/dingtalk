package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserDeleteRequest() *OapiUserDeleteRequest {
	return &OapiUserDeleteRequest{}
}

type OapiUserDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserDeleteResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiUserDeleteRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserDeleteRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiUserDeleteRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiUserDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.delete"
}
func (this *OapiUserDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiUserDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserDeleteResponse struct {
	taobao.TaobaoResponse
}
