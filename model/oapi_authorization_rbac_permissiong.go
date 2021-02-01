package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAuthorizationRbacPermissionGetRequest() *OapiAuthorizationRbacPermissionGetRequest {
	return &OapiAuthorizationRbacPermissionGetRequest{}
}

type OapiAuthorizationRbacPermissionGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthorizationRbacPermissionGetResponse
	AgentId         string
	Resource        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAuthorizationRbacPermissionGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthorizationRbacPermissionGetRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiAuthorizationRbacPermissionGetRequest) SetResource(resource2 string) {
	this.Resource = resource2
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetResource() string {
	return this.Resource
}
func (this *OapiAuthorizationRbacPermissionGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.authorization.rbac.permission.get"
}
func (this *OapiAuthorizationRbacPermissionGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthorizationRbacPermissionGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthorizationRbacPermissionGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["resource"] = this.Resource
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAuthorizationRbacPermissionGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthorizationRbacPermissionGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthorizationRbacPermissionGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenContactScopeVo struct {
	DeptIds                []int64  `json:"dept_ids,omitempty"`
	IncludeMemberDepts     bool     `json:"include_member_depts,omitempty"`
	IncludeSelfManageDepts bool     `json:"include_self_manage_depts,omitempty"`
	Userids                []string `json:"userids,omitempty"`
}
type OpenConditionVo struct {
	OpenContactScope OpenContactScopeVo `json:"open_contact_scope,omitempty"`
}
type OpenActionResultVo struct {
	ActionId      string          `json:"action_id,omitempty"`
	OpenCondition OpenConditionVo `json:"open_condition,omitempty"`
}
type PermitResultVo struct {
	OpenActionResults []OpenActionResultVo `json:"open_action_results,omitempty"`
	Permit            bool                 `json:"permit,omitempty"`
}
