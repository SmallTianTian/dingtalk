package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTasklistAddbyprojectRequest() *OapiWorkspaceTasklistAddbyprojectRequest {
	return &OapiWorkspaceTasklistAddbyprojectRequest{}
}

type OapiWorkspaceTasklistAddbyprojectRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTasklistAddbyprojectResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.tasklist.addbyproject"
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTasklistAddbyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTasklistAddbyprojectResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
