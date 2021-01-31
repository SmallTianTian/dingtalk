package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowancePredeductBetaRequest() *OapiCateringAllowancePredeductBetaRequest {
	return &OapiCateringAllowancePredeductBetaRequest{}
}

type OapiCateringAllowancePredeductBetaRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowancePredeductBetaResponse
	ExpiryTime      int64
	MealTime        int64
	OrderFullAmount int64
	Orderid         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringAllowancePredeductBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetExpiryTime(expiryTime2 int64) {
	this.ExpiryTime = expiryTime2
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetExpiryTime() int64 {
	return this.ExpiryTime
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetMealTime(mealTime2 int64) {
	this.MealTime = mealTime2
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetMealTime() int64 {
	return this.MealTime
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetOrderid(orderid2 string) {
	this.Orderid = orderid2
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetOrderid() string {
	return this.Orderid
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.prededuct.beta"
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowancePredeductBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["expiry_time"] = this.ExpiryTime
	txtParams["meal_time"] = this.MealTime
	txtParams["order_full_amount"] = this.OrderFullAmount
	txtParams["orderid"] = this.Orderid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringAllowancePredeductBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ExpiryTime, "expiryTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowancePredeductBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowancePredeductBetaResponse struct {
	taobao.TaobaoResponse
	Result CateringOpenPreDeductResponse `json:"result,omitempty"`
}
type CateringOpenPreDeductResponse struct {
	DeductedAmount int64  `json:"deducted_amount,omitempty"`
	MealPlanNo     string `json:"meal_plan_no,omitempty"`
}
