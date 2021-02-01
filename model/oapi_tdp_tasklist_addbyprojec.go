package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTasklistAddbyprojectRequest() *OapiTdpTasklistAddbyprojectRequest {
	return &OapiTdpTasklistAddbyprojectRequest{}
}

type OapiTdpTasklistAddbyprojectRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTasklistAddbyprojectResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiTdpTasklistAddbyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTasklistAddbyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTasklistAddbyprojectRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTasklistAddbyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpTasklistAddbyprojectRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.tasklist.addbyproject"
}
func (this *OapiTdpTasklistAddbyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTasklistAddbyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTasklistAddbyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiTdpTasklistAddbyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTasklistAddbyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTasklistAddbyprojectResponse struct {
	taobao.TaobaoResponse
}
