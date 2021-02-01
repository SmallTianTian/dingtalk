package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTasklistHidebyprojectRequest() *OapiWorkspaceTasklistHidebyprojectRequest {
	return &OapiWorkspaceTasklistHidebyprojectRequest{}
}

type OapiWorkspaceTasklistHidebyprojectRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTasklistHidebyprojectResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.tasklist.hidebyproject"
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTasklistHidebyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTasklistHidebyprojectResponse struct {
	taobao.TaobaoResponse
}
