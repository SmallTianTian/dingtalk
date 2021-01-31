package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingTaskCreateRequest() *OapiDingTaskCreateRequest {
	return &OapiDingTaskCreateRequest{}
}

type OapiDingTaskCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingTaskCreateResponse
	TaskSendVO      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingTaskCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingTaskCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingTaskCreateRequest) SetTaskSendVO(taskSendVO2 string) {
	this.TaskSendVO = taskSendVO2
}
func (this *OapiDingTaskCreateRequest) GetTaskSendVO() string {
	return this.TaskSendVO
}
func (this *OapiDingTaskCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ding.task.create"
}
func (this *OapiDingTaskCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingTaskCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingTaskCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingTaskCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingTaskCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["task_send_v_o"] = this.TaskSendVO
	return txtParams
}
func (this *OapiDingTaskCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingTaskCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingTaskCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingTaskCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  TaskSendResult `json:"result,omitempty"`
}
