package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordBatchupdateRequest() *OapiProcessWorkrecordBatchupdateRequest {
	return &OapiProcessWorkrecordBatchupdateRequest{}
}

type OapiProcessWorkrecordBatchupdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordBatchupdateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordBatchupdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordBatchupdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordBatchupdateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordBatchupdateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordBatchupdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.batchupdate"
}
func (this *OapiProcessWorkrecordBatchupdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordBatchupdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordBatchupdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordBatchupdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordBatchupdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordBatchupdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordBatchupdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordBatchupdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UpdateProcessInstanceRequest struct {
	ProcessInstanceId string `json:"process_instance_id,omitempty"`
	Result            string `json:"result,omitempty"`
	Status            string `json:"status,omitempty"`
}
type BatchUpdateProcessInstanceRequest struct {
	Agentid   int64                          `json:"agentid,omitempty"`
	Instances []UpdateProcessInstanceRequest `json:"instances,omitempty"`
}
type OapiProcessWorkrecordBatchupdateResponse struct {
	taobao.TaobaoResponse
}
