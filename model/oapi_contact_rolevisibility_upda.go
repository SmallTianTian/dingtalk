package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiContactRolevisibilityUpdateRequest() *OapiContactRolevisibilityUpdateRequest {
	return &OapiContactRolevisibilityUpdateRequest{}
}

type OapiContactRolevisibilityUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiContactRolevisibilityUpdateResponse
	Permissions     string
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiContactRolevisibilityUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiContactRolevisibilityUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiContactRolevisibilityUpdateRequest) SetPermissions(permissions2 string) {
	this.Permissions = permissions2
}
func (this *OapiContactRolevisibilityUpdateRequest) GetPermissions() string {
	return this.Permissions
}
func (this *OapiContactRolevisibilityUpdateRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiContactRolevisibilityUpdateRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiContactRolevisibilityUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.contact.rolevisibility.update"
}
func (this *OapiContactRolevisibilityUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiContactRolevisibilityUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiContactRolevisibilityUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiContactRolevisibilityUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiContactRolevisibilityUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["permissions"] = this.Permissions
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *OapiContactRolevisibilityUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Permissions, 999, "permissions"); err != nil {
		return err
	}
	return nil
}
func (this *OapiContactRolevisibilityUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiContactRolevisibilityUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiContactRolevisibilityUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
