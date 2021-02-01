package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappRuleGetRuleListRequest() *OapiMicroappRuleGetRuleListRequest {
	return &OapiMicroappRuleGetRuleListRequest{}
}

type OapiMicroappRuleGetRuleListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappRuleGetRuleListResponse
	AgentId         int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiMicroappRuleGetRuleListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappRuleGetRuleListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappRuleGetRuleListRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappRuleGetRuleListRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappRuleGetRuleListRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiMicroappRuleGetRuleListRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiMicroappRuleGetRuleListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.rule.get_rule_list"
}
func (this *OapiMicroappRuleGetRuleListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappRuleGetRuleListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappRuleGetRuleListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappRuleGetRuleListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappRuleGetRuleListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiMicroappRuleGetRuleListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappRuleGetRuleListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappRuleGetRuleListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappRuleGetRuleListResponse struct {
	taobao.TaobaoResponse

	RuleIdList []int64 `json:"ruleIdList,omitempty"`
}
