package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpMessageCorpconversationGetsendresultRequest() *CorpMessageCorpconversationGetsendresultRequest {
	return &CorpMessageCorpconversationGetsendresultRequest{}
}

type CorpMessageCorpconversationGetsendresultRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpMessageCorpconversationGetsendresultResponse
	AgentId         int64
	TaskId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpMessageCorpconversationGetsendresultRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpMessageCorpconversationGetsendresultRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *CorpMessageCorpconversationGetsendresultRequest) SetTaskId(taskId2 int64) {
	this.TaskId = taskId2
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetTaskId() int64 {
	return this.TaskId
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetApiMethodName() string {
	return "dingtalk.corp.message.corpconversation.getsendresult"
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpMessageCorpconversationGetsendresultRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpMessageCorpconversationGetsendresultRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpMessageCorpconversationGetsendresultRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *CorpMessageCorpconversationGetsendresultRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpMessageCorpconversationGetsendresultRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpMessageCorpconversationGetsendresultResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type AsyncSendResult struct {
	FailedUserIdList    []string `json:"failed_user_id_list,omitempty"`
	ForbiddenUserIdList []string `json:"forbidden_user_id_list,omitempty"`
	InvalidDeptIdList   []int64  `json:"invalid_dept_id_list,omitempty"`
	InvalidUserIdList   []string `json:"invalid_user_id_list,omitempty"`
	ReadUserIdList      []string `json:"read_user_id_list,omitempty"`
	UnreadUserIdList    []string `json:"unread_user_id_list,omitempty"`
}
