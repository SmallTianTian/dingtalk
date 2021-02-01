package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefActiveflowRequest() *OapiRhinoMosLayoutOperationdefActiveflowRequest {
	return &OapiRhinoMosLayoutOperationdefActiveflowRequest{}
}

type OapiRhinoMosLayoutOperationdefActiveflowRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosLayoutOperationdefActiveflowResponse
	FlowVersion     int64
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) SetFlowVersion(flowVersion2 int64) {
	this.FlowVersion = flowVersion2
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetFlowVersion() int64 {
	return this.FlowVersion
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdef.activeflow"
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["flow_version"] = this.FlowVersion
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FlowVersion, "flowVersion"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefActiveflowRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefActiveflowResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
