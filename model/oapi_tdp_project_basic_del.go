package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectBasicDeleteRequest() *OapiTdpProjectBasicDeleteRequest {
	return &OapiTdpProjectBasicDeleteRequest{}
}

type OapiTdpProjectBasicDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectBasicDeleteResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpProjectBasicDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectBasicDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectBasicDeleteRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectBasicDeleteRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectBasicDeleteRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectBasicDeleteRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectBasicDeleteRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectBasicDeleteRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectBasicDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.basic.delete"
}
func (this *OapiTdpProjectBasicDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectBasicDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectBasicDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectBasicDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectBasicDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	return txtParams
}
func (this *OapiTdpProjectBasicDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectBasicDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectBasicDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpProjectBasicDeleteResponse struct {
	taobao.TaobaoResponse
}
