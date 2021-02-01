package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleMemberAddRequest() *OapiAuthorizationRbacRoleMemberAddRequest {
	return &OapiAuthorizationRbacRoleMemberAddRequest{}
}

type OapiAuthorizationRbacRoleMemberAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleMemberAddResponse
	AddMembers      string
	AgentId         string
	OpenRoleId      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) SetAddMembers(addMembers2 string) {
	this.AddMembers = addMembers2
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetAddMembers() string {
	return this.AddMembers
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) SetOpenRoleId(openRoleId2 string) {
	this.OpenRoleId = openRoleId2
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetOpenRoleId() string {
	return this.OpenRoleId
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.member.add"
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["add_members"] = this.AddMembers
	txtParams["agent_id"] = this.AgentId
	txtParams["open_role_id"] = this.OpenRoleId
	return txtParams
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.AddMembers, 999, "addMembers"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleMemberAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacRoleMemberAddResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
