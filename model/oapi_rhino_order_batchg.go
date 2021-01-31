package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoOrderBatchGetRequest() *OapiRhinoOrderBatchGetRequest {
	return &OapiRhinoOrderBatchGetRequest{}
}

type OapiRhinoOrderBatchGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoOrderBatchGetResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoOrderBatchGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoOrderBatchGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoOrderBatchGetRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoOrderBatchGetRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoOrderBatchGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.order.batch.get"
}
func (this *OapiRhinoOrderBatchGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoOrderBatchGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoOrderBatchGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoOrderBatchGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoOrderBatchGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoOrderBatchGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoOrderBatchGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoOrderBatchGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiBatchGetProductOrderReq struct {
	IdList   []int64 `json:"id_list,omitempty"`
	TenantId string  `json:"tenant_id,omitempty"`
	Userid   string  `json:"userid,omitempty"`
}
type OapiRhinoOrderBatchGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                             `json:"errcode,omitempty"`
	Errmsg  string                            `json:"errmsg,omitempty"`
	Model   OpenApiBatchGetProductOrderResult `json:"model,omitempty"`
}
type OpenApiProductOrderDto struct {
	ActualFinishTime       time.Time `json:"actual_finish_time,omitempty"`
	ActualStartTime        time.Time `json:"actual_start_time,omitempty"`
	BizIdDtechCategory     string    `json:"biz_id_dtech_category,omitempty"`
	BizIdDtechCategoryName string    `json:"biz_id_dtech_category_name,omitempty"`
	BizIdDtechPkg          string    `json:"biz_id_dtech_pkg,omitempty"`
	ColorId                string    `json:"color_id,omitempty"`
	ColorName              string    `json:"color_name,omitempty"`
	GoodsNo                string    `json:"goods_no,omitempty"`
	Id                     int64     `json:"id,omitempty"`
	Number                 string    `json:"number,omitempty"`
	OrderCreateTime        time.Time `json:"order_create_time,omitempty"`
	PlanFinishTime         time.Time `json:"plan_finish_time,omitempty"`
	PlanStartTime          time.Time `json:"plan_start_time,omitempty"`
	PurchaseOrderId        string    `json:"purchase_order_id,omitempty"`
	Source                 string    `json:"source,omitempty"`
	Status                 string    `json:"status,omitempty"`
	StyleCode              string    `json:"style_code,omitempty"`
	StyleId                int64     `json:"style_id,omitempty"`
	StyleImg               string    `json:"style_img,omitempty"`
	StyleName              string    `json:"style_name,omitempty"`
	TenantId               string    `json:"tenant_id,omitempty"`
	TotalQuantity          int64     `json:"total_quantity,omitempty"`
}
type OpenApiBatchGetProductOrderResult struct {
	List []OpenApiProductOrderDto `json:"list,omitempty"`
}
