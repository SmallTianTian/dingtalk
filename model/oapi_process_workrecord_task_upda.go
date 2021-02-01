package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordTaskUpdateRequest() *OapiProcessWorkrecordTaskUpdateRequest {
	return &OapiProcessWorkrecordTaskUpdateRequest{}
}

type OapiProcessWorkrecordTaskUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordTaskUpdateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordTaskUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.task.update"
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordTaskUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UpdateTaskRequest struct {
	Agentid           int64       `json:"agentid,omitempty"`
	ProcessInstanceId string      `json:"process_instance_id,omitempty"`
	Tasks             []TaskTopVo `json:"tasks,omitempty"`
}
type OapiProcessWorkrecordTaskUpdateResponse struct {
	taobao.TaobaoResponse
}
