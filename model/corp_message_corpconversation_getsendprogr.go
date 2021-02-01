package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpMessageCorpconversationGetsendprogressRequest() *CorpMessageCorpconversationGetsendprogressRequest {
	return &CorpMessageCorpconversationGetsendprogressRequest{}
}

type CorpMessageCorpconversationGetsendprogressRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpMessageCorpconversationGetsendprogressResponse
	AgentId         int64
	TaskId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpMessageCorpconversationGetsendprogressRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) SetTaskId(taskId2 int64) {
	this.TaskId = taskId2
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetTaskId() int64 {
	return this.TaskId
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetApiMethodName() string {
	return "dingtalk.corp.message.corpconversation.getsendprogress"
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpMessageCorpconversationGetsendprogressRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpMessageCorpconversationGetsendprogressResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type AsyncSendProgress struct {
	ProgressInPercent int64 `json:"progress_in_percent,omitempty"`
	Status            int64 `json:"status,omitempty"`
}
