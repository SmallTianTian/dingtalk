package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTasklistHidebyprojectRequest() *OapiTdpTasklistHidebyprojectRequest {
	return &OapiTdpTasklistHidebyprojectRequest{}
}

type OapiTdpTasklistHidebyprojectRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTasklistHidebyprojectResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiTdpTasklistHidebyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTasklistHidebyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTasklistHidebyprojectRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTasklistHidebyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpTasklistHidebyprojectRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.tasklist.hidebyproject"
}
func (this *OapiTdpTasklistHidebyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTasklistHidebyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTasklistHidebyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiTdpTasklistHidebyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTasklistHidebyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTasklistHidebyprojectResponse struct {
	taobao.TaobaoResponse
}
