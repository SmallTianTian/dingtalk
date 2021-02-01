package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefsEditassignRequest() *OapiRhinoMosLayoutOperationdefsEditassignRequest {
	return &OapiRhinoMosLayoutOperationdefsEditassignRequest{}
}

type OapiRhinoMosLayoutOperationdefsEditassignRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiRhinoMosLayoutOperationdefsEditassignResponse
	AssignInfoModifyItems string
	OrderId               int64
	TenantId              string
	TopHttpMethod         string
	TopResponseType       string
	Userid                string
}

func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) SetAssignInfoModifyItems(assignInfoModifyItems2 string) {
	this.AssignInfoModifyItems = assignInfoModifyItems2
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetAssignInfoModifyItems() string {
	return this.AssignInfoModifyItems
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdefs.editassign"
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["assign_info_modify_items"] = this.AssignInfoModifyItems
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.AssignInfoModifyItems, 500, "assignInfoModifyItems"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefsEditassignRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AssignInfoModifyItem struct {
	OperationExecutorAssigns []OperationExecutorAssignDto `json:"operation_executor_assigns,omitempty"`
	OperationExternalId      string                       `json:"operation_external_id,omitempty"`
	OperationUid             int64                        `json:"operation_uid,omitempty"`
	WorkUnits                []WorkUnitDto                `json:"work_units,omitempty"`
}
type OapiRhinoMosLayoutOperationdefsEditassignResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
