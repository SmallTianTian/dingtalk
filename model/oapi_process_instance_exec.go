package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessInstanceExecuteRequest() *OapiProcessInstanceExecuteRequest {
	return &OapiProcessInstanceExecuteRequest{}
}

type OapiProcessInstanceExecuteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessInstanceExecuteResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessInstanceExecuteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessInstanceExecuteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessInstanceExecuteRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessInstanceExecuteRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessInstanceExecuteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.instance.execute"
}
func (this *OapiProcessInstanceExecuteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessInstanceExecuteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessInstanceExecuteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessInstanceExecuteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessInstanceExecuteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessInstanceExecuteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessInstanceExecuteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessInstanceExecuteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ExecuteTaskRequest struct {
	ActionerUserid    string `json:"actioner_userid,omitempty"`
	File              File   `json:"file,omitempty"`
	ProcessInstanceId string `json:"process_instance_id,omitempty"`
	Remark            string `json:"remark,omitempty"`
	Result            string `json:"result,omitempty"`
	TaskId            int64  `json:"task_id,omitempty"`
}
type OapiProcessInstanceExecuteResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
