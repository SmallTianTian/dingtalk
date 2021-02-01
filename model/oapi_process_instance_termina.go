package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessInstanceTerminateRequest() *OapiProcessInstanceTerminateRequest {
	return &OapiProcessInstanceTerminateRequest{}
}

type OapiProcessInstanceTerminateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessInstanceTerminateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessInstanceTerminateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessInstanceTerminateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessInstanceTerminateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessInstanceTerminateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessInstanceTerminateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.instance.terminate"
}
func (this *OapiProcessInstanceTerminateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessInstanceTerminateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessInstanceTerminateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessInstanceTerminateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessInstanceTerminateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessInstanceTerminateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessInstanceTerminateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessInstanceTerminateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TerminateProcessInstanceRequestV2 struct {
	IsSystem          bool   `json:"is_system,omitempty"`
	OperatingUserid   string `json:"operating_userid,omitempty"`
	ProcessInstanceId string `json:"process_instance_id,omitempty"`
	Remark            string `json:"remark,omitempty"`
}
type OapiProcessInstanceTerminateResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
