package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiContactRolevisibilityGetRequest() *OapiContactRolevisibilityGetRequest {
	return &OapiContactRolevisibilityGetRequest{}
}

type OapiContactRolevisibilityGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiContactRolevisibilityGetResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiContactRolevisibilityGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiContactRolevisibilityGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiContactRolevisibilityGetRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiContactRolevisibilityGetRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiContactRolevisibilityGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.contact.rolevisibility.get"
}
func (this *OapiContactRolevisibilityGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiContactRolevisibilityGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiContactRolevisibilityGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiContactRolevisibilityGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiContactRolevisibilityGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *OapiContactRolevisibilityGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiContactRolevisibilityGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiContactRolevisibilityGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiContactRolevisibilityGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                          `json:"errcode,omitempty"`
	Errmsg  string                         `json:"errmsg,omitempty"`
	Result  []OpenRolePermissionForContact `json:"result,omitempty"`
	Success bool                           `json:"success,omitempty"`
}
type OpenRolePermissionForContact struct {
	DeptIds        []int64  `json:"dept_ids,omitempty"`
	RoleIds        []int64  `json:"role_ids,omitempty"`
	Side           int64    `json:"side,omitempty"`
	Type           int64    `json:"type,omitempty"`
	UserIds        []string `json:"user_ids,omitempty"`
	VisibilityType int64    `json:"visibility_type,omitempty"`
}
