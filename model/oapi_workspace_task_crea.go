package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskCreateRequest() *OapiWorkspaceTaskCreateRequest {
	return &OapiWorkspaceTaskCreateRequest{}
}

type OapiWorkspaceTaskCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskCreateResponse
	MicroappAgentId int64
	OperatorUserid  string
	Task            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskCreateRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTaskCreateRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTaskCreateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTaskCreateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTaskCreateRequest) SetTask(task2 string) {
	this.Task = task2
}
func (this *OapiWorkspaceTaskCreateRequest) GetTask() string {
	return this.Task
}
func (this *OapiWorkspaceTaskCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.create"
}
func (this *OapiWorkspaceTaskCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task"] = this.Task
	return txtParams
}
func (this *OapiWorkspaceTaskCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TaskSystemProperty struct {
	AppUrl  string `json:"appUrl,omitempty"`
	GroupId string `json:"groupId,omitempty"`
	WebUrl  string `json:"webUrl,omitempty"`
}
type OapiWorkspaceTaskCreateResponse struct {
	taobao.TaobaoResponse
	Result Task `json:"result,omitempty"`
}
