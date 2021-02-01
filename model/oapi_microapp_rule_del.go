package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappRuleDeleteRequest() *OapiMicroappRuleDeleteRequest {
	return &OapiMicroappRuleDeleteRequest{}
}

type OapiMicroappRuleDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappRuleDeleteResponse
	AgentId         int64
	RuleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappRuleDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappRuleDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappRuleDeleteRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappRuleDeleteRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappRuleDeleteRequest) SetRuleId(ruleId2 int64) {
	this.RuleId = ruleId2
}
func (this *OapiMicroappRuleDeleteRequest) GetRuleId() int64 {
	return this.RuleId
}
func (this *OapiMicroappRuleDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.rule.delete"
}
func (this *OapiMicroappRuleDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappRuleDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappRuleDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappRuleDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappRuleDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams["ruleId"] = this.RuleId
	return txtParams
}
func (this *OapiMicroappRuleDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappRuleDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappRuleDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappRuleDeleteResponse struct {
	taobao.TaobaoResponse
}
