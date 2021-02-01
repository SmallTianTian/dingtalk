package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingTaskStatusUpdateRequest() *OapiDingTaskStatusUpdateRequest {
	return &OapiDingTaskStatusUpdateRequest{}
}

type OapiDingTaskStatusUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingTaskStatusUpdateResponse
	OperatorUserid  string
	TaskId          string
	TaskStatus      int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingTaskStatusUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingTaskStatusUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingTaskStatusUpdateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiDingTaskStatusUpdateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiDingTaskStatusUpdateRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiDingTaskStatusUpdateRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiDingTaskStatusUpdateRequest) SetTaskStatus(taskStatus2 int64) {
	this.TaskStatus = taskStatus2
}
func (this *OapiDingTaskStatusUpdateRequest) GetTaskStatus() int64 {
	return this.TaskStatus
}
func (this *OapiDingTaskStatusUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ding.task.status.update"
}
func (this *OapiDingTaskStatusUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingTaskStatusUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingTaskStatusUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingTaskStatusUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingTaskStatusUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task_id"] = this.TaskId
	txtParams["task_status"] = this.TaskStatus
	return txtParams
}
func (this *OapiDingTaskStatusUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingTaskStatusUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingTaskStatusUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingTaskStatusUpdateResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
