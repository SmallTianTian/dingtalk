package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappSetVisibleScopesRequest() *OapiMicroappSetVisibleScopesRequest {
	return &OapiMicroappSetVisibleScopesRequest{}
}

type OapiMicroappSetVisibleScopesRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiMicroappSetVisibleScopesResponse
	AgentId           int64
	DeptVisibleScopes []int64
	IsHidden          bool
	TopHttpMethod     string
	TopResponseType   string
	UserVisibleScopes []string
}

func (this *OapiMicroappSetVisibleScopesRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappSetVisibleScopesRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappSetVisibleScopesRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappSetVisibleScopesRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappSetVisibleScopesRequest) SetDeptVisibleScopes(deptVisibleScopes2 []int64) {
	this.DeptVisibleScopes = deptVisibleScopes2
}
func (this *OapiMicroappSetVisibleScopesRequest) GetDeptVisibleScopes() []int64 {
	return this.DeptVisibleScopes
}
func (this *OapiMicroappSetVisibleScopesRequest) SetIsHidden(isHidden2 bool) {
	this.IsHidden = isHidden2
}
func (this *OapiMicroappSetVisibleScopesRequest) GetIsHidden() bool {
	return this.IsHidden
}
func (this *OapiMicroappSetVisibleScopesRequest) SetUserVisibleScopes(userVisibleScopes2 []string) {
	this.UserVisibleScopes = userVisibleScopes2
}
func (this *OapiMicroappSetVisibleScopesRequest) GetUserVisibleScopes() []string {
	return this.UserVisibleScopes
}
func (this *OapiMicroappSetVisibleScopesRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.set_visible_scopes"
}
func (this *OapiMicroappSetVisibleScopesRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappSetVisibleScopesRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappSetVisibleScopesRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappSetVisibleScopesRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappSetVisibleScopesRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams["deptVisibleScopes"] = this.DeptVisibleScopes
	txtParams["isHidden"] = this.IsHidden
	txtParams["userVisibleScopes"] = this.UserVisibleScopes
	return txtParams
}
func (this *OapiMicroappSetVisibleScopesRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappSetVisibleScopesRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappSetVisibleScopesRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappSetVisibleScopesResponse struct {
	taobao.TaobaoResponse
}
