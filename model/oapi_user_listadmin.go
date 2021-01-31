package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserListadminRequest() *OapiUserListadminRequest {
	return &OapiUserListadminRequest{}
}

type OapiUserListadminRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserListadminResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserListadminRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserListadminRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserListadminRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.listadmin"
}
func (this *OapiUserListadminRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserListadminRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserListadminRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserListadminRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserListadminRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiUserListadminRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserListadminRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserListadminRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserListadminResponse struct {
	taobao.TaobaoResponse
	Result []ListAdminResponse `json:"result,omitempty"`
}
type ListAdminResponse struct {
	SysLevel int64  `json:"sys_level,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
