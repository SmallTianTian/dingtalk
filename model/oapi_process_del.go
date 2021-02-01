package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessDeleteRequest() *OapiProcessDeleteRequest {
	return &OapiProcessDeleteRequest{}
}

type OapiProcessDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessDeleteResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessDeleteRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessDeleteRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.delete"
}
func (this *OapiProcessDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeleteProcessRequest struct {
	Agentid          int64  `json:"agentid,omitempty"`
	CleanRunningTask bool   `json:"clean_running_task,omitempty"`
	ProcessCode      string `json:"process_code,omitempty"`
}
type OapiProcessDeleteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
