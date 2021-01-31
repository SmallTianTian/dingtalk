package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTaskBasicDeletebyprojectRequest() *OapiTdpTaskBasicDeletebyprojectRequest {
	return &OapiTdpTaskBasicDeletebyprojectRequest{}
}

type OapiTdpTaskBasicDeletebyprojectRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTaskBasicDeletebyprojectResponse
	MicroappAgentId int64
	OperatorUserid  string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.task.basic.deletebyproject"
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project_id"] = this.ProjectId
	return txtParams
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTaskBasicDeletebyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTaskBasicDeletebyprojectResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
