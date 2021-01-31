package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleListRequest() *OapiAuthorizationRbacRoleListRequest {
	return &OapiAuthorizationRbacRoleListRequest{}
}

type OapiAuthorizationRbacRoleListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleListResponse
	AgentId         string
	Cursor          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleListRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleListRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiAuthorizationRbacRoleListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiAuthorizationRbacRoleListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAuthorizationRbacRoleListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAuthorizationRbacRoleListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.list"
}
func (this *OapiAuthorizationRbacRoleListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAuthorizationRbacRoleListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacRoleListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenEamRoleVo struct {
	OpenAction    OpenActionVo   `json:"open_action,omitempty"`
	OpenMembers   []OpenMemberVo `json:"open_members,omitempty"`
	OpenResources []string       `json:"open_resources,omitempty"`
	OpenRoleId    string         `json:"open_role_id,omitempty"`
	OpenRoleName  string         `json:"open_role_name,omitempty"`
}
