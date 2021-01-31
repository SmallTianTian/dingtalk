package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleRemoveRequest() *OapiAuthorizationRbacRoleRemoveRequest {
	return &OapiAuthorizationRbacRoleRemoveRequest{}
}

type OapiAuthorizationRbacRoleRemoveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleRemoveResponse
	AgentId         string
	OpenRoleId      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) SetOpenRoleId(openRoleId2 string) {
	this.OpenRoleId = openRoleId2
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetOpenRoleId() string {
	return this.OpenRoleId
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.remove"
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["open_role_id"] = this.OpenRoleId
	return txtParams
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacRoleRemoveResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
