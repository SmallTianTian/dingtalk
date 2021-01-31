package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordTaskCreateRequest() *OapiProcessWorkrecordTaskCreateRequest {
	return &OapiProcessWorkrecordTaskCreateRequest{}
}

type OapiProcessWorkrecordTaskCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordTaskCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordTaskCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordTaskCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordTaskCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordTaskCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordTaskCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.task.create"
}
func (this *OapiProcessWorkrecordTaskCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordTaskCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordTaskCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordTaskCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordTaskCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordTaskCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordTaskCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordTaskCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TaskTopVo struct {
	Url    string `json:"url,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type SaveTaskRequest struct {
	ActivityId        string      `json:"activity_id,omitempty"`
	Agentid           int64       `json:"agentid,omitempty"`
	ProcessInstanceId string      `json:"process_instance_id,omitempty"`
	Tasks             []TaskTopVo `json:"tasks,omitempty"`
}
type OapiProcessWorkrecordTaskCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64       `json:"errcode,omitempty"`
	Errmsg  string      `json:"errmsg,omitempty"`
	Tasks   []TaskTopVo `json:"tasks,omitempty"`
}
