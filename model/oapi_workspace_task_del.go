package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskDeleteRequest() *OapiWorkspaceTaskDeleteRequest {
	return &OapiWorkspaceTaskDeleteRequest{}
}

type OapiWorkspaceTaskDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskDeleteResponse
	MicroappAgentId int64
	OperatorUserid  string
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskDeleteRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTaskDeleteRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTaskDeleteRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTaskDeleteRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTaskDeleteRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiWorkspaceTaskDeleteRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiWorkspaceTaskDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.delete"
}
func (this *OapiWorkspaceTaskDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiWorkspaceTaskDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTaskDeleteResponse struct {
	taobao.TaobaoResponse
}
