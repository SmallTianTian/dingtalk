package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleResourceUpdateRequest() *OapiAuthorizationRbacRoleResourceUpdateRequest {
	return &OapiAuthorizationRbacRoleResourceUpdateRequest{}
}

type OapiAuthorizationRbacRoleResourceUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleResourceUpdateResponse
	AgentId         string
	OpenResources   string
	OpenRoleId      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) SetOpenResources(openResources2 string) {
	this.OpenResources = openResources2
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetOpenResources() string {
	return this.OpenResources
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) SetOpenRoleId(openRoleId2 string) {
	this.OpenRoleId = openRoleId2
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetOpenRoleId() string {
	return this.OpenRoleId
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.resource.update"
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["open_resources"] = this.OpenResources
	txtParams["open_role_id"] = this.OpenRoleId
	return txtParams
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleResourceUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacRoleResourceUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
