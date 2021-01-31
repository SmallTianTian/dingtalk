package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectMemberBatchremoveRequest() *OapiTdpProjectMemberBatchremoveRequest {
	return &OapiTdpProjectMemberBatchremoveRequest{}
}

type OapiTdpProjectMemberBatchremoveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectMemberBatchremoveResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiTdpProjectMemberBatchremoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectMemberBatchremoveRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectMemberBatchremoveRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectMemberBatchremoveRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectMemberBatchremoveRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.member.batchremove"
}
func (this *OapiTdpProjectMemberBatchremoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectMemberBatchremoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectMemberBatchremoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiTdpProjectMemberBatchremoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectMemberBatchremoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpProjectMemberBatchremoveResponse struct {
	taobao.TaobaoResponse
}
