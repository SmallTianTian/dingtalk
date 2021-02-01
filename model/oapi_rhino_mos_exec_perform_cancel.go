package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformCancelRequest() *OapiRhinoMosExecPerformCancelRequest {
	return &OapiRhinoMosExecPerformCancelRequest{}
}

type OapiRhinoMosExecPerformCancelRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                      OapiRhinoMosExecPerformCancelResponse
	Context                   string
	OperationPerformRecordIds string
	OrderId                   int64
	StopSchedule              bool
	TenantId                  string
	TopHttpMethod             string
	TopResponseType           string
	Userid                    string
}

func (this *OapiRhinoMosExecPerformCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetContext(context2 string) {
	this.Context = context2
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetContextString(context2 string) {
	this.Context = context2
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetContext() string {
	return this.Context
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetOperationPerformRecordIds(operationPerformRecordIds2 string) {
	this.OperationPerformRecordIds = operationPerformRecordIds2
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetOperationPerformRecordIds() string {
	return this.OperationPerformRecordIds
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetStopSchedule(stopSchedule2 bool) {
	this.StopSchedule = stopSchedule2
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetStopSchedule() bool {
	return this.StopSchedule
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.cancel"
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["context"] = this.Context
	txtParams["operation_perform_record_ids"] = this.OperationPerformRecordIds
	txtParams["order_id"] = this.OrderId
	txtParams["stop_schedule"] = this.StopSchedule
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecPerformCancelRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperationPerformRecordIds, "operationPerformRecordIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformCancelResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
