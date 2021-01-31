package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskUpdateRequest() *OapiWorkspaceTaskUpdateRequest {
	return &OapiWorkspaceTaskUpdateRequest{}
}

type OapiWorkspaceTaskUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskUpdateResponse
	MicroappAgentId int64
	OperatorUserid  string
	Task            string
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskUpdateRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTaskUpdateRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTaskUpdateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTaskUpdateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTaskUpdateRequest) SetTask(task2 string) {
	this.Task = task2
}
func (this *OapiWorkspaceTaskUpdateRequest) GetTask() string {
	return this.Task
}
func (this *OapiWorkspaceTaskUpdateRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiWorkspaceTaskUpdateRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiWorkspaceTaskUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.update"
}
func (this *OapiWorkspaceTaskUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task"] = this.Task
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiWorkspaceTaskUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskId, "taskId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTaskUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
