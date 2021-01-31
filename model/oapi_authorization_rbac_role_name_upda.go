package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleNameUpdateRequest() *OapiAuthorizationRbacRoleNameUpdateRequest {
	return &OapiAuthorizationRbacRoleNameUpdateRequest{}
}

type OapiAuthorizationRbacRoleNameUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleNameUpdateResponse
	AgentId         string
	OpenRoleId      string
	RoleName        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) SetOpenRoleId(openRoleId2 string) {
	this.OpenRoleId = openRoleId2
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetOpenRoleId() string {
	return this.OpenRoleId
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) SetRoleName(roleName2 string) {
	this.RoleName = roleName2
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetRoleName() string {
	return this.RoleName
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.name.update"
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["open_role_id"] = this.OpenRoleId
	txtParams["role_name"] = this.RoleName
	return txtParams
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleNameUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacRoleNameUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
