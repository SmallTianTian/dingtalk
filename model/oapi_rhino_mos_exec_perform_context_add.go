package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformContextAddRequest() *OapiRhinoMosExecPerformContextAddRequest {
	return &OapiRhinoMosExecPerformContextAddRequest{}
}

type OapiRhinoMosExecPerformContextAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiRhinoMosExecPerformContextAddResponse
	Context            string
	OperationRecordIds string
	OrderId            string
	TenantId           string
	TopHttpMethod      string
	TopResponseType    string
	Userid             string
}

func (this *OapiRhinoMosExecPerformContextAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetContext(context2 string) {
	this.Context = context2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetContextString(context2 string) {
	this.Context = context2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetContext() string {
	return this.Context
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetOperationRecordIds(operationRecordIds2 string) {
	this.OperationRecordIds = operationRecordIds2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetOperationRecordIds() string {
	return this.OperationRecordIds
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.context.add"
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformContextAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["context"] = this.Context
	txtParams["operation_record_ids"] = this.OperationRecordIds
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecPerformContextAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Context, "context"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformContextAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformContextAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64                 `json:"errcode,omitempty"`
	Errmsg  string                `json:"errmsg,omitempty"`
	Model   []OperationPerformDto `json:"model,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
