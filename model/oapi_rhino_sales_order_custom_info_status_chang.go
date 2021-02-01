package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoSalesOrderCustomInfoStatusChangeRequest() *OapiRhinoSalesOrderCustomInfoStatusChangeRequest {
	return &OapiRhinoSalesOrderCustomInfoStatusChangeRequest{}
}

type OapiRhinoSalesOrderCustomInfoStatusChangeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                          OapiRhinoSalesOrderCustomInfoStatusChangeResponse
	SalesOrderCustomInfoChangeReq string
	TopHttpMethod                 string
	TopResponseType               string
}

func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) SetSalesOrderCustomInfoChangeReq(salesOrderCustomInfoChangeReq2 string) {
	this.SalesOrderCustomInfoChangeReq = salesOrderCustomInfoChangeReq2
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetSalesOrderCustomInfoChangeReq() string {
	return this.SalesOrderCustomInfoChangeReq
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.sales.order.custom.info.status.change"
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["sales_order_custom_info_change_req"] = this.SalesOrderCustomInfoChangeReq
	return txtParams
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoSalesOrderCustomInfoStatusChangeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiCustomOrderChangeReq struct {
	BizIdCustomerOrder string `json:"biz_id_customer_order,omitempty"`
	Status             string `json:"status,omitempty"`
	TenantId           string `json:"tenant_id,omitempty"`
	Userid             string `json:"userid,omitempty"`
}
type OapiRhinoSalesOrderCustomInfoStatusChangeResponse struct {
	taobao.TaobaoResponse

	ExternalMsgInfo string `json:"external_msg_info,omitempty"`
	Model           bool   `json:"model,omitempty"`
	Success         bool   `json:"success,omitempty"`
}
