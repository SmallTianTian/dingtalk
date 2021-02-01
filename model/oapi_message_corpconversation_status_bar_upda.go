package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageCorpconversationStatusBarUpdateRequest() *OapiMessageCorpconversationStatusBarUpdateRequest {
	return &OapiMessageCorpconversationStatusBarUpdateRequest{}
}

type OapiMessageCorpconversationStatusBarUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationStatusBarUpdateResponse
	AgentId         int64
	StatusBg        string
	StatusValue     string
	TaskId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) SetStatusBg(statusBg2 string) {
	this.StatusBg = statusBg2
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetStatusBg() string {
	return this.StatusBg
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) SetStatusValue(statusValue2 string) {
	this.StatusValue = statusValue2
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetStatusValue() string {
	return this.StatusValue
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) SetTaskId(taskId2 int64) {
	this.TaskId = taskId2
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetTaskId() int64 {
	return this.TaskId
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.status_bar.update"
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["status_bg"] = this.StatusBg
	txtParams["status_value"] = this.StatusValue
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationStatusBarUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageCorpconversationStatusBarUpdateResponse struct {
	taobao.TaobaoResponse
}
