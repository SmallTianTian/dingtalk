package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreOrdersSpecialCanalUpdateOrderRequest() *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest {
	return &OapiAppstoreOrdersSpecialCanalUpdateOrderRequest{}
}

type OapiAppstoreOrdersSpecialCanalUpdateOrderRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreOrdersSpecialCanalUpdateOrderResponse
	DingOrderId     string
	Status          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) SetDingOrderId(dingOrderId2 string) {
	this.DingOrderId = dingOrderId2
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetDingOrderId() string {
	return this.DingOrderId
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.orders.special_canal.update_order"
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ding_order_id"] = this.DingOrderId
	txtParams["status"] = this.Status
	return txtParams
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DingOrderId, "dingOrderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreOrdersSpecialCanalUpdateOrderRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreOrdersSpecialCanalUpdateOrderResponse struct {
	taobao.TaobaoResponse
}
