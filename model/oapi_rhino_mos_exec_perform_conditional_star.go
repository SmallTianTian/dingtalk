package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformConditionalStartRequest() *OapiRhinoMosExecPerformConditionalStartRequest {
	return &OapiRhinoMosExecPerformConditionalStartRequest{}
}

type OapiRhinoMosExecPerformConditionalStartRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecPerformConditionalStartResponse
	DeviceIds       string
	EntityCondition string
	OperationUids   string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WorkNos         string
}

func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetDeviceIds(deviceIds2 string) {
	this.DeviceIds = deviceIds2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetDeviceIds() string {
	return this.DeviceIds
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetEntityCondition(entityCondition2 string) {
	this.EntityCondition = entityCondition2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetEntityCondition() string {
	return this.EntityCondition
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetOperationUids(operationUids2 string) {
	this.OperationUids = operationUids2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetOperationUids() string {
	return this.OperationUids
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetWorkNos(workNos2 string) {
	this.WorkNos = workNos2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetWorkNos() string {
	return this.WorkNos
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.conditional.start"
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_ids"] = this.DeviceIds
	txtParams["entity_condition"] = this.EntityCondition
	txtParams["operation_uids"] = this.OperationUids
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_nos"] = this.WorkNos
	return txtParams
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeviceIds, 20, "deviceIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformConditionalStartRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformConditionalStartResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
