package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiContactRolevisibilityDeleteRequest() *OapiContactRolevisibilityDeleteRequest {
	return &OapiContactRolevisibilityDeleteRequest{}
}

type OapiContactRolevisibilityDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiContactRolevisibilityDeleteResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiContactRolevisibilityDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiContactRolevisibilityDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiContactRolevisibilityDeleteRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiContactRolevisibilityDeleteRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiContactRolevisibilityDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.contact.rolevisibility.delete"
}
func (this *OapiContactRolevisibilityDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiContactRolevisibilityDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiContactRolevisibilityDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiContactRolevisibilityDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiContactRolevisibilityDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *OapiContactRolevisibilityDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiContactRolevisibilityDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiContactRolevisibilityDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiContactRolevisibilityDeleteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
