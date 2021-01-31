package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreInternalUnfinishedorderListRequest() *OapiAppstoreInternalUnfinishedorderListRequest {
	return &OapiAppstoreInternalUnfinishedorderListRequest{}
}

type OapiAppstoreInternalUnfinishedorderListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreInternalUnfinishedorderListResponse
	ItemCode        string
	Page            int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) SetItemCode(itemCode2 string) {
	this.ItemCode = itemCode2
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetItemCode() string {
	return this.ItemCode
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) SetPage(page2 int64) {
	this.Page = page2
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetPage() int64 {
	return this.Page
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.internal.unfinishedorder.list"
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["item_code"] = this.ItemCode
	txtParams["page"] = this.Page
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Page, "page"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreInternalUnfinishedorderListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreInternalUnfinishedorderListResponse struct {
	taobao.TaobaoResponse
	Result PageModel `json:"result,omitempty"`
}
type InAppGoodsOrderVO struct {
	BizOrderId        int64  `json:"biz_order_id,omitempty"`
	CorpId            string `json:"corp_id,omitempty"`
	CreateTimestamp   int64  `json:"create_timestamp,omitempty"`
	GoodsCode         string `json:"goods_code,omitempty"`
	ItemCode          string `json:"item_code,omitempty"`
	PaidTimestamp     int64  `json:"paid_timestamp,omitempty"`
	Quantity          int64  `json:"quantity,omitempty"`
	Status            int64  `json:"status,omitempty"`
	TotalActualPayFee int64  `json:"total_actual_pay_fee,omitempty"`
}
type PageModel struct {
	Items []InAppGoodsOrderVO `json:"items,omitempty"`
	Total int64               `json:"total,omitempty"`
}
