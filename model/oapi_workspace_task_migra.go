package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskMigrateRequest() *OapiWorkspaceTaskMigrateRequest {
	return &OapiWorkspaceTaskMigrateRequest{}
}

type OapiWorkspaceTaskMigrateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskMigrateResponse
	Agentid         int64
	OperatorUserid  string
	Task            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskMigrateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskMigrateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskMigrateRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWorkspaceTaskMigrateRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWorkspaceTaskMigrateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTaskMigrateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTaskMigrateRequest) SetTask(task2 string) {
	this.Task = task2
}
func (this *OapiWorkspaceTaskMigrateRequest) GetTask() string {
	return this.Task
}
func (this *OapiWorkspaceTaskMigrateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.migrate"
}
func (this *OapiWorkspaceTaskMigrateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskMigrateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskMigrateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskMigrateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskMigrateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task"] = this.Task
	return txtParams
}
func (this *OapiWorkspaceTaskMigrateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskMigrateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskMigrateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTaskMigrateResponse struct {
	taobao.TaobaoResponse
	Result Task `json:"result,omitempty"`
}
