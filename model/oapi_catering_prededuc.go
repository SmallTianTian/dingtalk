package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringPredeductRequest() *OapiCateringPredeductRequest {
	return &OapiCateringPredeductRequest{}
}

type OapiCateringPredeductRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringPredeductResponse
	ExpiryTime      int64
	MealTime        int64
	OrderFullAmount int64
	Orderid         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringPredeductRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringPredeductRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringPredeductRequest) SetExpiryTime(expiryTime2 int64) {
	this.ExpiryTime = expiryTime2
}
func (this *OapiCateringPredeductRequest) GetExpiryTime() int64 {
	return this.ExpiryTime
}
func (this *OapiCateringPredeductRequest) SetMealTime(mealTime2 int64) {
	this.MealTime = mealTime2
}
func (this *OapiCateringPredeductRequest) GetMealTime() int64 {
	return this.MealTime
}
func (this *OapiCateringPredeductRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringPredeductRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringPredeductRequest) SetOrderid(orderid2 string) {
	this.Orderid = orderid2
}
func (this *OapiCateringPredeductRequest) GetOrderid() string {
	return this.Orderid
}
func (this *OapiCateringPredeductRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringPredeductRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringPredeductRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.prededuct"
}
func (this *OapiCateringPredeductRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringPredeductRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringPredeductRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringPredeductRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringPredeductRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["expiry_time"] = this.ExpiryTime
	txtParams["meal_time"] = this.MealTime
	txtParams["order_full_amount"] = this.OrderFullAmount
	txtParams["orderid"] = this.Orderid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringPredeductRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ExpiryTime, "expiryTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringPredeductRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringPredeductRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringPredeductResponse struct {
	taobao.TaobaoResponse
	Result CateringOpenPreDeductResponse `json:"result,omitempty"`
}
