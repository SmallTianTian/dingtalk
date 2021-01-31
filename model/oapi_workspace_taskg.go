package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskGetRequest() *OapiWorkspaceTaskGetRequest {
	return &OapiWorkspaceTaskGetRequest{}
}

type OapiWorkspaceTaskGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskGetResponse
	MicroappAgentId int64
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskGetRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTaskGetRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTaskGetRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiWorkspaceTaskGetRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiWorkspaceTaskGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.get"
}
func (this *OapiWorkspaceTaskGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiWorkspaceTaskGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskId, "taskId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTaskGetResponse struct {
	taobao.TaobaoResponse
	Result Task `json:"result,omitempty"`
}
