package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappVisibleScopesRequest() *OapiMicroappVisibleScopesRequest {
	return &OapiMicroappVisibleScopesRequest{}
}

type OapiMicroappVisibleScopesRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappVisibleScopesResponse
	AgentId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappVisibleScopesRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappVisibleScopesRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappVisibleScopesRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappVisibleScopesRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappVisibleScopesRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.visible_scopes"
}
func (this *OapiMicroappVisibleScopesRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappVisibleScopesRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappVisibleScopesRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappVisibleScopesRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappVisibleScopesRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	return txtParams
}
func (this *OapiMicroappVisibleScopesRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappVisibleScopesRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappVisibleScopesRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappVisibleScopesResponse struct {
	taobao.TaobaoResponse
	DeptVisibleScopes []int64 `json:"deptVisibleScopes,omitempty"`

	IsHidden          bool     `json:"isHidden,omitempty"`
	UserVisibleScopes []string `json:"userVisibleScopes,omitempty"`
}
