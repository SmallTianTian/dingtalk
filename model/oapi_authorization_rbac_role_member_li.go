package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleMemberListRequest() *OapiAuthorizationRbacRoleMemberListRequest {
	return &OapiAuthorizationRbacRoleMemberListRequest{}
}

type OapiAuthorizationRbacRoleMemberListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleMemberListResponse
	AgentId         string
	Cursor          int64
	OpenRoleId      string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleMemberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) SetOpenRoleId(openRoleId2 string) {
	this.OpenRoleId = openRoleId2
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetOpenRoleId() string {
	return this.OpenRoleId
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.member.list"
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["cursor"] = this.Cursor
	txtParams["open_role_id"] = this.OpenRoleId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleMemberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacRoleMemberListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
