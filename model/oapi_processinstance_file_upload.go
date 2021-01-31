package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceFileUploadRequest() *OapiProcessinstanceFileUploadRequest {
	return &OapiProcessinstanceFileUploadRequest{}
}

type OapiProcessinstanceFileUploadRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiProcessinstanceFileUploadResponse
	AgentId           int64
	ProcessInstanceId string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiProcessinstanceFileUploadRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceFileUploadRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceFileUploadRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiProcessinstanceFileUploadRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiProcessinstanceFileUploadRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *OapiProcessinstanceFileUploadRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *OapiProcessinstanceFileUploadRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.file.upload"
}
func (this *OapiProcessinstanceFileUploadRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceFileUploadRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceFileUploadRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceFileUploadRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceFileUploadRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["process_instance_id"] = this.ProcessInstanceId
	return txtParams
}
func (this *OapiProcessinstanceFileUploadRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceFileUploadRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceFileUploadRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessinstanceFileUploadResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  AppSpaceResponse `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
