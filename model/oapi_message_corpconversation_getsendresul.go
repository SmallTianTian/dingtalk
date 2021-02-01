package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMessageCorpconversationGetsendresultRequest() *OapiMessageCorpconversationGetsendresultRequest {
	return &OapiMessageCorpconversationGetsendresultRequest{}
}

type OapiMessageCorpconversationGetsendresultRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationGetsendresultResponse
	AgentId         int64
	TaskId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMessageCorpconversationGetsendresultRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationGetsendresultRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationGetsendresultRequest) SetTaskId(taskId2 int64) {
	this.TaskId = taskId2
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetTaskId() int64 {
	return this.TaskId
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.getsendresult"
}
func (this *OapiMessageCorpconversationGetsendresultRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationGetsendresultRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationGetsendresultRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiMessageCorpconversationGetsendresultRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationGetsendresultRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageCorpconversationGetsendresultResponse struct {
	taobao.TaobaoResponse

	SendResult AsyncSendResult `json:"send_result,omitempty"`
}
type SendForbiddenModel struct {
	Code   string `json:"code,omitempty"`
	Count  int64  `json:"count,omitempty"`
	Userid string `json:"userid,omitempty"`
}
