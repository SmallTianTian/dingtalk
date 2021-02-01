package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCspaceGetCustomSpaceRequest() *OapiCspaceGetCustomSpaceRequest {
	return &OapiCspaceGetCustomSpaceRequest{}
}

type OapiCspaceGetCustomSpaceRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceGetCustomSpaceResponse
	AgentId         string
	Domain          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCspaceGetCustomSpaceRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiCspaceGetCustomSpaceRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceGetCustomSpaceRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiCspaceGetCustomSpaceRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiCspaceGetCustomSpaceRequest) SetDomain(domain2 string) {
	this.Domain = domain2
}
func (this *OapiCspaceGetCustomSpaceRequest) GetDomain() string {
	return this.Domain
}
func (this *OapiCspaceGetCustomSpaceRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.get_custom_space"
}
func (this *OapiCspaceGetCustomSpaceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceGetCustomSpaceRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceGetCustomSpaceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceGetCustomSpaceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceGetCustomSpaceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["domain"] = this.Domain
	return txtParams
}
func (this *OapiCspaceGetCustomSpaceRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCspaceGetCustomSpaceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceGetCustomSpaceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceGetCustomSpaceResponse struct {
	taobao.TaobaoResponse
	Spaceid string `json:"spaceid,omitempty"`
}
