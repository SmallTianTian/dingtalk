package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTaskBasicDeleteRequest() *OapiTdpTaskBasicDeleteRequest {
	return &OapiTdpTaskBasicDeleteRequest{}
}

type OapiTdpTaskBasicDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTaskBasicDeleteResponse
	MicroappAgentId int64
	OperatorUserid  string
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpTaskBasicDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTaskBasicDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTaskBasicDeleteRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTaskBasicDeleteRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTaskBasicDeleteRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTaskBasicDeleteRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTaskBasicDeleteRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiTdpTaskBasicDeleteRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiTdpTaskBasicDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.task.basic.delete"
}
func (this *OapiTdpTaskBasicDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTaskBasicDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTaskBasicDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTaskBasicDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTaskBasicDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiTdpTaskBasicDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTaskBasicDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTaskBasicDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTaskBasicDeleteResponse struct {
	taobao.TaobaoResponse
}
