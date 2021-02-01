package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefsListsimpleRequest() *OapiRhinoMosLayoutOperationdefsListsimpleRequest {
	return &OapiRhinoMosLayoutOperationdefsListsimpleRequest{}
}

type OapiRhinoMosLayoutOperationdefsListsimpleRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosLayoutOperationdefsListsimpleResponse
	OperationUids   string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) SetOperationUids(operationUids2 string) {
	this.OperationUids = operationUids2
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetOperationUids() string {
	return this.OperationUids
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdefs.listsimple"
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operation_uids"] = this.OperationUids
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperationUids, "operationUids"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefsListsimpleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefsListsimpleResponse struct {
	taobao.TaobaoResponse
	Result []OperationDefDto `json:"result,omitempty"`
}
