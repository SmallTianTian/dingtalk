package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectMemberRemovebyprojectRequest() *OapiTdpProjectMemberRemovebyprojectRequest {
	return &OapiTdpProjectMemberRemovebyprojectRequest{}
}

type OapiTdpProjectMemberRemovebyprojectRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectMemberRemovebyprojectResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.member.removebyproject"
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	return txtParams
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectMemberRemovebyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpProjectMemberRemovebyprojectResponse struct {
	taobao.TaobaoResponse
}
