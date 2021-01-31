package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTaskBasicGetRequest() *OapiTdpTaskBasicGetRequest {
	return &OapiTdpTaskBasicGetRequest{}
}

type OapiTdpTaskBasicGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTaskBasicGetResponse
	MicroappAgentId int64
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpTaskBasicGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTaskBasicGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTaskBasicGetRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTaskBasicGetRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTaskBasicGetRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiTdpTaskBasicGetRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiTdpTaskBasicGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.task.basic.get"
}
func (this *OapiTdpTaskBasicGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTaskBasicGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTaskBasicGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTaskBasicGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTaskBasicGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiTdpTaskBasicGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskId, "taskId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTaskBasicGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTaskBasicGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTaskBasicGetResponse struct {
	taobao.TaobaoResponse
	Result OrgTask `json:"result,omitempty"`
}
