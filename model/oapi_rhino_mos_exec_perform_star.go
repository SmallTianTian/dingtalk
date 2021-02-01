package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformStartRequest() *OapiRhinoMosExecPerformStartRequest {
	return &OapiRhinoMosExecPerformStartRequest{}
}

type OapiRhinoMosExecPerformStartRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                      OapiRhinoMosExecPerformStartResponse
	DeviceIds                 string
	OperationPerformRecordIds string
	OrderId                   int64
	TenantId                  string
	TopHttpMethod             string
	TopResponseType           string
	Userid                    string
	WorkNos                   string
}

func (this *OapiRhinoMosExecPerformStartRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformStartRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformStartRequest) SetDeviceIds(deviceIds2 string) {
	this.DeviceIds = deviceIds2
}
func (this *OapiRhinoMosExecPerformStartRequest) GetDeviceIds() string {
	return this.DeviceIds
}
func (this *OapiRhinoMosExecPerformStartRequest) SetOperationPerformRecordIds(operationPerformRecordIds2 string) {
	this.OperationPerformRecordIds = operationPerformRecordIds2
}
func (this *OapiRhinoMosExecPerformStartRequest) GetOperationPerformRecordIds() string {
	return this.OperationPerformRecordIds
}
func (this *OapiRhinoMosExecPerformStartRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformStartRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformStartRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformStartRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformStartRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformStartRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformStartRequest) SetWorkNos(workNos2 string) {
	this.WorkNos = workNos2
}
func (this *OapiRhinoMosExecPerformStartRequest) GetWorkNos() string {
	return this.WorkNos
}
func (this *OapiRhinoMosExecPerformStartRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.start"
}
func (this *OapiRhinoMosExecPerformStartRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformStartRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformStartRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformStartRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformStartRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_ids"] = this.DeviceIds
	txtParams["operation_perform_record_ids"] = this.OperationPerformRecordIds
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_nos"] = this.WorkNos
	return txtParams
}
func (this *OapiRhinoMosExecPerformStartRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeviceIds, 500, "deviceIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformStartRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformStartRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformStartResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
