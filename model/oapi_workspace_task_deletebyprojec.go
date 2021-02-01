package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskDeletebyprojectRequest() *OapiWorkspaceTaskDeletebyprojectRequest {
	return &OapiWorkspaceTaskDeletebyprojectRequest{}
}

type OapiWorkspaceTaskDeletebyprojectRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskDeletebyprojectResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.deletebyproject"
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	return txtParams
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskDeletebyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTaskDeletebyprojectResponse struct {
	taobao.TaobaoResponse
}
