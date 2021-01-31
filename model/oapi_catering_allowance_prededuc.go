package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowancePredeductRequest() *OapiCateringAllowancePredeductRequest {
	return &OapiCateringAllowancePredeductRequest{}
}

type OapiCateringAllowancePredeductRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowancePredeductResponse
	ExpiryTime      int64
	MealTime        int64
	OrderFullAmount int64
	Orderid         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringAllowancePredeductRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowancePredeductRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowancePredeductRequest) SetExpiryTime(expiryTime2 int64) {
	this.ExpiryTime = expiryTime2
}
func (this *OapiCateringAllowancePredeductRequest) GetExpiryTime() int64 {
	return this.ExpiryTime
}
func (this *OapiCateringAllowancePredeductRequest) SetMealTime(mealTime2 int64) {
	this.MealTime = mealTime2
}
func (this *OapiCateringAllowancePredeductRequest) GetMealTime() int64 {
	return this.MealTime
}
func (this *OapiCateringAllowancePredeductRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringAllowancePredeductRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringAllowancePredeductRequest) SetOrderid(orderid2 string) {
	this.Orderid = orderid2
}
func (this *OapiCateringAllowancePredeductRequest) GetOrderid() string {
	return this.Orderid
}
func (this *OapiCateringAllowancePredeductRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowancePredeductRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowancePredeductRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.prededuct"
}
func (this *OapiCateringAllowancePredeductRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowancePredeductRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowancePredeductRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowancePredeductRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowancePredeductRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["expiry_time"] = this.ExpiryTime
	txtParams["meal_time"] = this.MealTime
	txtParams["order_full_amount"] = this.OrderFullAmount
	txtParams["orderid"] = this.Orderid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringAllowancePredeductRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ExpiryTime, "expiryTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowancePredeductRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowancePredeductRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowancePredeductResponse struct {
	taobao.TaobaoResponse
	Errcode int64                         `json:"errcode,omitempty"`
	Errmsg  string                        `json:"errmsg,omitempty"`
	Result  CateringOpenPreDeductResponse `json:"result,omitempty"`
}
