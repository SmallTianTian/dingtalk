package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleMemberRemoveRequest() *OapiAuthorizationRbacRoleMemberRemoveRequest {
	return &OapiAuthorizationRbacRoleMemberRemoveRequest{}
}

type OapiAuthorizationRbacRoleMemberRemoveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleMemberRemoveResponse
	AgentId         string
	OpenRoleId      string
	RemoveMembers   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) SetOpenRoleId(openRoleId2 string) {
	this.OpenRoleId = openRoleId2
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetOpenRoleId() string {
	return this.OpenRoleId
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) SetRemoveMembers(removeMembers2 string) {
	this.RemoveMembers = removeMembers2
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetRemoveMembers() string {
	return this.RemoveMembers
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.member.remove"
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["open_role_id"] = this.OpenRoleId
	txtParams["remove_members"] = this.RemoveMembers
	return txtParams
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleMemberRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacRoleMemberRemoveResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
