package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecOperationConditionInactiveRequest() *OapiRhinoMosExecOperationConditionInactiveRequest {
	return &OapiRhinoMosExecOperationConditionInactiveRequest{}
}

type OapiRhinoMosExecOperationConditionInactiveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiRhinoMosExecOperationConditionInactiveResponse
	InactiveOperationReq string
	TopHttpMethod        string
	TopResponseType      string
}

func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) SetInactiveOperationReq(inactiveOperationReq2 string) {
	this.InactiveOperationReq = inactiveOperationReq2
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetInactiveOperationReq() string {
	return this.InactiveOperationReq
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.operation.condition.inactive"
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["inactive_operation_req"] = this.InactiveOperationReq
	return txtParams
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecOperationConditionInactiveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OperationCondition struct {
	Ids    []int64 `json:"ids,omitempty"`
	Source Source  `json:"source,omitempty"`
}
type InactiveOperationByConditionReq struct {
	Condition OperationCondition `json:"condition,omitempty"`
	OrderId   int64              `json:"order_id,omitempty"`
	TenantId  string             `json:"tenant_id,omitempty"`
	Userid    string             `json:"userid,omitempty"`
}
type OapiRhinoMosExecOperationConditionInactiveResponse struct {
	taobao.TaobaoResponse
	Errcode         int64  `json:"errcode,omitempty"`
	Errmsg          string `json:"errmsg,omitempty"`
	ExternalMsgInfo string `json:"external_msg_info,omitempty"`
	Model           bool   `json:"model,omitempty"`
}
