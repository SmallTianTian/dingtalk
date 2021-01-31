package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleUpdateRoleRequest() *OapiRoleUpdateRoleRequest {
	return &OapiRoleUpdateRoleRequest{}
}

type OapiRoleUpdateRoleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleUpdateRoleResponse
	RoleId          int64
	RoleName        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleUpdateRoleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleUpdateRoleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleUpdateRoleRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleUpdateRoleRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleUpdateRoleRequest) SetRoleName(roleName2 string) {
	this.RoleName = roleName2
}
func (this *OapiRoleUpdateRoleRequest) GetRoleName() string {
	return this.RoleName
}
func (this *OapiRoleUpdateRoleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.update_role"
}
func (this *OapiRoleUpdateRoleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleUpdateRoleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleUpdateRoleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleUpdateRoleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleUpdateRoleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["roleId"] = this.RoleId
	txtParams["roleName"] = this.RoleName
	return txtParams
}
func (this *OapiRoleUpdateRoleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleUpdateRoleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleUpdateRoleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleUpdateRoleResponse struct {
	taobao.TaobaoResponse
}
