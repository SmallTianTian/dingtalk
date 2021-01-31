package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformFinishRequest() *OapiRhinoMosExecPerformFinishRequest {
	return &OapiRhinoMosExecPerformFinishRequest{}
}

type OapiRhinoMosExecPerformFinishRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                      OapiRhinoMosExecPerformFinishResponse
	DeviceIds                 string
	OperationPerformRecordIds string
	OrderId                   int64
	ProcessCostTime           string
	TenantId                  string
	TopHttpMethod             string
	TopResponseType           string
	Userid                    string
	WorkNos                   string
}

func (this *OapiRhinoMosExecPerformFinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetDeviceIds(deviceIds2 string) {
	this.DeviceIds = deviceIds2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetDeviceIds() string {
	return this.DeviceIds
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetOperationPerformRecordIds(operationPerformRecordIds2 string) {
	this.OperationPerformRecordIds = operationPerformRecordIds2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetOperationPerformRecordIds() string {
	return this.OperationPerformRecordIds
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetProcessCostTime(processCostTime2 string) {
	this.ProcessCostTime = processCostTime2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetProcessCostTime() string {
	return this.ProcessCostTime
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetWorkNos(workNos2 string) {
	this.WorkNos = workNos2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetWorkNos() string {
	return this.WorkNos
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.finish"
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformFinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_ids"] = this.DeviceIds
	txtParams["operation_perform_record_ids"] = this.OperationPerformRecordIds
	txtParams["order_id"] = this.OrderId
	txtParams["process_cost_time"] = this.ProcessCostTime
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_nos"] = this.WorkNos
	return txtParams
}
func (this *OapiRhinoMosExecPerformFinishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeviceIds, 20, "deviceIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformFinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformFinishResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
