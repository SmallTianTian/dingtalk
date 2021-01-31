package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringDeductCapacityRequest() *OapiCateringDeductCapacityRequest {
	return &OapiCateringDeductCapacityRequest{}
}

type OapiCateringDeductCapacityRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringDeductCapacityResponse
	MeaTime         int64
	OrderFullAmount int64
	Orderid         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringDeductCapacityRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringDeductCapacityRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringDeductCapacityRequest) SetMeaTime(meaTime2 int64) {
	this.MeaTime = meaTime2
}
func (this *OapiCateringDeductCapacityRequest) GetMeaTime() int64 {
	return this.MeaTime
}
func (this *OapiCateringDeductCapacityRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringDeductCapacityRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringDeductCapacityRequest) SetOrderid(orderid2 string) {
	this.Orderid = orderid2
}
func (this *OapiCateringDeductCapacityRequest) GetOrderid() string {
	return this.Orderid
}
func (this *OapiCateringDeductCapacityRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringDeductCapacityRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringDeductCapacityRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.deduct.capacity"
}
func (this *OapiCateringDeductCapacityRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringDeductCapacityRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringDeductCapacityRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringDeductCapacityRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringDeductCapacityRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mea_time"] = this.MeaTime
	txtParams["order_full_amount"] = this.OrderFullAmount
	txtParams["orderid"] = this.Orderid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringDeductCapacityRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MeaTime, "meaTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringDeductCapacityRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringDeductCapacityRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringDeductCapacityResponse struct {
	taobao.TaobaoResponse
	Result CateringOpenGetDeductCapacityResponse `json:"result,omitempty"`
}
