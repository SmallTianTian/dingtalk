package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreInternalOrderGetRequest() *OapiAppstoreInternalOrderGetRequest {
	return &OapiAppstoreInternalOrderGetRequest{}
}

type OapiAppstoreInternalOrderGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreInternalOrderGetResponse
	BizOrderId      int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreInternalOrderGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreInternalOrderGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreInternalOrderGetRequest) SetBizOrderId(bizOrderId2 int64) {
	this.BizOrderId = bizOrderId2
}
func (this *OapiAppstoreInternalOrderGetRequest) GetBizOrderId() int64 {
	return this.BizOrderId
}
func (this *OapiAppstoreInternalOrderGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.internal.order.get"
}
func (this *OapiAppstoreInternalOrderGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreInternalOrderGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreInternalOrderGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreInternalOrderGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreInternalOrderGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_order_id"] = this.BizOrderId
	return txtParams
}
func (this *OapiAppstoreInternalOrderGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizOrderId, "bizOrderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreInternalOrderGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreInternalOrderGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreInternalOrderGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  InAppGoodsOrderVo `json:"result,omitempty"`
}
type InAppGoodsOrderVo struct {
	BizOrderId        int64  `json:"biz_order_id,omitempty"`
	CorpId            string `json:"corp_id,omitempty"`
	CreateTimestamp   int64  `json:"create_timestamp,omitempty"`
	EndTimestamp      int64  `json:"end_timestamp,omitempty"`
	GoodsCode         string `json:"goods_code,omitempty"`
	ItemCode          string `json:"item_code,omitempty"`
	PaidTimestamp     int64  `json:"paid_timestamp,omitempty"`
	Quantity          int64  `json:"quantity,omitempty"`
	StartTimestamp    int64  `json:"start_timestamp,omitempty"`
	Status            int64  `json:"status,omitempty"`
	TotalActualPayFee int64  `json:"total_actual_pay_fee,omitempty"`
}
