package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformConditionalFinishRequest() *OapiRhinoMosExecPerformConditionalFinishRequest {
	return &OapiRhinoMosExecPerformConditionalFinishRequest{}
}

type OapiRhinoMosExecPerformConditionalFinishRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecPerformConditionalFinishResponse
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

func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetDeviceIds(deviceIds2 string) {
	this.DeviceIds = deviceIds2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetDeviceIds() string {
	return this.DeviceIds
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetEntityCondition(entityCondition2 string) {
	this.EntityCondition = entityCondition2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetEntityCondition() string {
	return this.EntityCondition
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetOperationUids(operationUids2 string) {
	this.OperationUids = operationUids2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetOperationUids() string {
	return this.OperationUids
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetWorkNos(workNos2 string) {
	this.WorkNos = workNos2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetWorkNos() string {
	return this.WorkNos
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.conditional.finish"
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeviceIds, 20, "deviceIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformConditionalFinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformConditionalFinishResponse struct {
	taobao.TaobaoResponse
	Errcode         int64  `json:"errcode,omitempty"`
	Errmsg          string `json:"errmsg,omitempty"`
	ExternalMsgInfo string `json:"external_msg_info,omitempty"`
	Model           bool   `json:"model,omitempty"`
	Success         bool   `json:"success,omitempty"`
}
