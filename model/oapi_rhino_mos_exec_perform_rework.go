package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformReworkRequest() *OapiRhinoMosExecPerformReworkRequest {
	return &OapiRhinoMosExecPerformReworkRequest{}
}

type OapiRhinoMosExecPerformReworkRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecPerformReworkResponse
	Context         string
	OrderId         int64
	ReworkStartId   int64
	TenantId        string
	ToInactiveIds   string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosExecPerformReworkRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetContext(context2 string) {
	this.Context = context2
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetContextString(context2 string) {
	this.Context = context2
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetContext() string {
	return this.Context
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetReworkStartId(reworkStartId2 int64) {
	this.ReworkStartId = reworkStartId2
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetReworkStartId() int64 {
	return this.ReworkStartId
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetToInactiveIds(toInactiveIds2 string) {
	this.ToInactiveIds = toInactiveIds2
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetToInactiveIds() string {
	return this.ToInactiveIds
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.rework"
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformReworkRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["context"] = this.Context
	txtParams["order_id"] = this.OrderId
	txtParams["rework_start_id"] = this.ReworkStartId
	txtParams["tenant_id"] = this.TenantId
	txtParams["to_inactive_ids"] = this.ToInactiveIds
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecPerformReworkRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReworkStartId, "reworkStartId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformReworkRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformReworkResponse struct {
	taobao.TaobaoResponse
	Model   OperationPerformDto `json:"model,omitempty"`
	Success bool                `json:"success,omitempty"`
}
