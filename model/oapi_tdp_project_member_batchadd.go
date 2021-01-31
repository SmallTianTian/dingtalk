package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectMemberBatchaddRequest() *OapiTdpProjectMemberBatchaddRequest {
	return &OapiTdpProjectMemberBatchaddRequest{}
}

type OapiTdpProjectMemberBatchaddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectMemberBatchaddResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiTdpProjectMemberBatchaddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectMemberBatchaddRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectMemberBatchaddRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectMemberBatchaddRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectMemberBatchaddRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.member.batchadd"
}
func (this *OapiTdpProjectMemberBatchaddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectMemberBatchaddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectMemberBatchaddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiTdpProjectMemberBatchaddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectMemberBatchaddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpProjectMemberBatchaddResponse struct {
	taobao.TaobaoResponse
}
