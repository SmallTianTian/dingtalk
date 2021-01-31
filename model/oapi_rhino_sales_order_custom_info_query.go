package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoSalesOrderCustomInfoQueryRequest() *OapiRhinoSalesOrderCustomInfoQueryRequest {
	return &OapiRhinoSalesOrderCustomInfoQueryRequest{}
}

type OapiRhinoSalesOrderCustomInfoQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoSalesOrderCustomInfoQueryResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.sales.order.custom.info.query"
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoSalesOrderCustomInfoQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiSalesOrderCustomInfoQueryReq struct {
	BatchId                  int64     `json:"batch_id,omitempty"`
	Page                     int64     `json:"page,omitempty"`
	PageSize                 int64     `json:"page_size,omitempty"`
	PlannedDeliveryTimeBegin time.Time `json:"planned_delivery_time_begin,omitempty"`
	PlannedDeliveryTimeEnd   time.Time `json:"planned_delivery_time_end,omitempty"`
	ProductOrderId           int64     `json:"product_order_id,omitempty"`
	TenantId                 string    `json:"tenant_id,omitempty"`
	Userid                   string    `json:"userid,omitempty"`
}
type OapiRhinoSalesOrderCustomInfoQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64                                  `json:"errcode,omitempty"`
	Errmsg  string                                 `json:"errmsg,omitempty"`
	Model   OpenApiSalesOrderQueryCustomInfoResult `json:"model,omitempty"`
}
type OpenApiSalesOrderCustomInfoDto struct {
	BizIdCustomerOrder           string    `json:"biz_id_customer_order,omitempty"`
	GmtOrderCreate               time.Time `json:"gmt_order_create,omitempty"`
	GmtPlannedDelivery           time.Time `json:"gmt_planned_delivery,omitempty"`
	GmtPlannedProductionFinished time.Time `json:"gmt_planned_production_finished,omitempty"`
	ImgUrl                       string    `json:"img_url,omitempty"`
	ProductOrderId               int64     `json:"product_order_id,omitempty"`
	ProductionTime               time.Time `json:"production_time,omitempty"`
	Quantity                     int64     `json:"quantity,omitempty"`
	SizeId                       int64     `json:"size_id,omitempty"`
	SizeName                     string    `json:"size_name,omitempty"`
	SkipCustomized               bool      `json:"skip_customized,omitempty"`
	Status                       string    `json:"status,omitempty"`
}
type OpenApiSalesOrderQueryCustomInfoResult struct {
	List     []OpenApiSalesOrderCustomInfoDto `json:"list,omitempty"`
	Page     int64                            `json:"page,omitempty"`
	PageSize int64                            `json:"page_size,omitempty"`
	Total    int64                            `json:"total,omitempty"`
}
