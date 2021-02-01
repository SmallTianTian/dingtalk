package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageCorpconversationGetsendprogressRequest() *OapiMessageCorpconversationGetsendprogressRequest {
	return &OapiMessageCorpconversationGetsendprogressRequest{}
}

type OapiMessageCorpconversationGetsendprogressRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationGetsendprogressResponse
	AgentId         int64
	TaskId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMessageCorpconversationGetsendprogressRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) SetTaskId(taskId2 int64) {
	this.TaskId = taskId2
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetTaskId() int64 {
	return this.TaskId
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.getsendprogress"
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationGetsendprogressRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageCorpconversationGetsendprogressResponse struct {
	taobao.TaobaoResponse

	Progress AsyncSendProgress `json:"progress,omitempty"`
}
