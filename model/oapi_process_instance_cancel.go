package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessInstanceCancelRequest() *OapiProcessInstanceCancelRequest {
	return &OapiProcessInstanceCancelRequest{}
}

type OapiProcessInstanceCancelRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessInstanceCancelResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessInstanceCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessInstanceCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessInstanceCancelRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessInstanceCancelRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessInstanceCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.instance.cancel"
}
func (this *OapiProcessInstanceCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessInstanceCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessInstanceCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessInstanceCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessInstanceCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessInstanceCancelRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessInstanceCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessInstanceCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CancelProcessInstanceRequest struct {
	OperatorName      string `json:"operator_name,omitempty"`
	ProcessInstanceId string `json:"process_instance_id,omitempty"`
	Remark            string `json:"remark,omitempty"`
}
type OapiProcessInstanceCancelResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
