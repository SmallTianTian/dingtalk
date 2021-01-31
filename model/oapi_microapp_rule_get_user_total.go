package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappRuleGetUserTotalRequest() *OapiMicroappRuleGetUserTotalRequest {
	return &OapiMicroappRuleGetUserTotalRequest{}
}

type OapiMicroappRuleGetUserTotalRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappRuleGetUserTotalResponse
	AgentId         int64
	RuleIdList      []int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappRuleGetUserTotalRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappRuleGetUserTotalRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappRuleGetUserTotalRequest) SetRuleIdList(ruleIdList2 []int64) {
	this.RuleIdList = ruleIdList2
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetRuleIdList() []int64 {
	return this.RuleIdList
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.rule.get_user_total"
}
func (this *OapiMicroappRuleGetUserTotalRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappRuleGetUserTotalRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappRuleGetUserTotalRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams["ruleIdList"] = this.RuleIdList
	return txtParams
}
func (this *OapiMicroappRuleGetUserTotalRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappRuleGetUserTotalRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappRuleGetUserTotalResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Result  []Result `json:"result,omitempty"`
}
