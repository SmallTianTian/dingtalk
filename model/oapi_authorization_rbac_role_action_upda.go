package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleActionUpdateRequest() *OapiAuthorizationRbacRoleActionUpdateRequest {
	return &OapiAuthorizationRbacRoleActionUpdateRequest{}
}

type OapiAuthorizationRbacRoleActionUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleActionUpdateResponse
	AgentId         string
	OpenAction      string
	OpenRoleId      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) SetOpenAction(openAction2 string) {
	this.OpenAction = openAction2
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetOpenAction() string {
	return this.OpenAction
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) SetOpenRoleId(openRoleId2 string) {
	this.OpenRoleId = openRoleId2
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetOpenRoleId() string {
	return this.OpenRoleId
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.action.update"
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["open_action"] = this.OpenAction
	txtParams["open_role_id"] = this.OpenRoleId
	return txtParams
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleActionUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenActionVo struct {
	ActionIds     []string        `json:"action_ids,omitempty"`
	OpenCondition OpenConditionVo `json:"open_condition,omitempty"`
}
type OapiAuthorizationRbacRoleActionUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
