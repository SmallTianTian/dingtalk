package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacRoleCreateRequest() *OapiAuthorizationRbacRoleCreateRequest {
	return &OapiAuthorizationRbacRoleCreateRequest{}
}

type OapiAuthorizationRbacRoleCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacRoleCreateResponse
	AgentId         string
	OpenRoleCreate  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthorizationRbacRoleCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacRoleCreateRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacRoleCreateRequest) SetOpenRoleCreate(openRoleCreate2 string) {
	this.OpenRoleCreate = openRoleCreate2
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetOpenRoleCreate() string {
	return this.OpenRoleCreate
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.role.create"
}
func (this *OapiAuthorizationRbacRoleCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacRoleCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacRoleCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["open_role_create"] = this.OpenRoleCreate
	return txtParams
}
func (this *OapiAuthorizationRbacRoleCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacRoleCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenMemberVo struct {
	BelongCorpid  string `json:"belong_corpid,omitempty"`
	MemberId      string `json:"member_id,omitempty"`
	MemberType    string `json:"member_type,omitempty"`
	OperateUserid string `json:"operate_userid,omitempty"`
}
type OpenRoleCreateVo struct {
	OpenAction    OpenActionVo   `json:"open_action,omitempty"`
	OpenMembers   []OpenMemberVo `json:"open_members,omitempty"`
	OpenResources []string       `json:"open_resources,omitempty"`
	OpenRoleId    string         `json:"open_role_id,omitempty"`
	OpenRoleName  string         `json:"open_role_name,omitempty"`
}
type OapiAuthorizationRbacRoleCreateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
