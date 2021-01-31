package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectMemberAddRequest() *OapiTdpProjectMemberAddRequest {
	return &OapiTdpProjectMemberAddRequest{}
}

type OapiTdpProjectMemberAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectMemberAddResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiTdpProjectMemberAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectMemberAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectMemberAddRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectMemberAddRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectMemberAddRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectMemberAddRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectMemberAddRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectMemberAddRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectMemberAddRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiTdpProjectMemberAddRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiTdpProjectMemberAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.member.add"
}
func (this *OapiTdpProjectMemberAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectMemberAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectMemberAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectMemberAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectMemberAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiTdpProjectMemberAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectMemberAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectMemberAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpProjectMemberAddResponse struct {
	taobao.TaobaoResponse
	Result Project `json:"result,omitempty"`
}
