package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceFileDownloadRequest() *OapiProcessinstanceFileDownloadRequest {
	return &OapiProcessinstanceFileDownloadRequest{}
}

type OapiProcessinstanceFileDownloadRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp              OapiProcessinstanceFileDownloadResponse
	AgentId           int64
	FileId            string
	ProcessInstanceId string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiProcessinstanceFileDownloadRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceFileDownloadRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceFileDownloadRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiProcessinstanceFileDownloadRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiProcessinstanceFileDownloadRequest) SetFileId(fileId2 string) {
	this.FileId = fileId2
}
func (this *OapiProcessinstanceFileDownloadRequest) GetFileId() string {
	return this.FileId
}
func (this *OapiProcessinstanceFileDownloadRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *OapiProcessinstanceFileDownloadRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *OapiProcessinstanceFileDownloadRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.file.download"
}
func (this *OapiProcessinstanceFileDownloadRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceFileDownloadRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceFileDownloadRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceFileDownloadRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceFileDownloadRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["file_id"] = this.FileId
	txtParams["process_instance_id"] = this.ProcessInstanceId
	return txtParams
}
func (this *OapiProcessinstanceFileDownloadRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceFileDownloadRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceFileDownloadRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessinstanceFileDownloadResponse struct {
	taobao.TaobaoResponse
	Result  AppSpaceResponse `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
