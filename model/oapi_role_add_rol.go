package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleAddRoleRequest() *OapiRoleAddRoleRequest {
	return &OapiRoleAddRoleRequest{}
}

type OapiRoleAddRoleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleAddRoleResponse
	GroupId         int64
	RoleName        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleAddRoleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleAddRoleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleAddRoleRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiRoleAddRoleRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiRoleAddRoleRequest) SetRoleName(roleName2 string) {
	this.RoleName = roleName2
}
func (this *OapiRoleAddRoleRequest) GetRoleName() string {
	return this.RoleName
}
func (this *OapiRoleAddRoleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.add_role"
}
func (this *OapiRoleAddRoleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleAddRoleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleAddRoleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleAddRoleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleAddRoleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["groupId"] = this.GroupId
	txtParams["roleName"] = this.RoleName
	return txtParams
}
func (this *OapiRoleAddRoleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleAddRoleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleAddRoleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleAddRoleResponse struct {
	taobao.TaobaoResponse
	RoleId int64 `json:"roleId,omitempty"`
}
