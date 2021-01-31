package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowanceOrderPushBetaRequest() *OapiCateringAllowanceOrderPushBetaRequest {
	return &OapiCateringAllowanceOrderPushBetaRequest{}
}

type OapiCateringAllowanceOrderPushBetaRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowanceOrderPushBetaResponse
	ActualAmount    int64
	AllowanceAmount int64
	Ext             string
	MealPlanNo      string
	MealTime        int64
	OrderDetails    string
	OrderFullAmount int64
	OrderId         string
	OrderTime       int64
	ShopId          string
	ShopName        string
	TopHttpMethod   string
	TopResponseType string
	UserName        string
	Userid          string
}

func (this *OapiCateringAllowanceOrderPushBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetActualAmount(actualAmount2 int64) {
	this.ActualAmount = actualAmount2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetActualAmount() int64 {
	return this.ActualAmount
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetAllowanceAmount(allowanceAmount2 int64) {
	this.AllowanceAmount = allowanceAmount2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetAllowanceAmount() int64 {
	return this.AllowanceAmount
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetExt(ext2 string) {
	this.Ext = ext2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetExt() string {
	return this.Ext
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetMealPlanNo(mealPlanNo2 string) {
	this.MealPlanNo = mealPlanNo2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetMealPlanNo() string {
	return this.MealPlanNo
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetMealTime(mealTime2 int64) {
	this.MealTime = mealTime2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetMealTime() int64 {
	return this.MealTime
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetOrderDetails(orderDetails2 string) {
	this.OrderDetails = orderDetails2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetOrderDetails() string {
	return this.OrderDetails
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetOrderTime(orderTime2 int64) {
	this.OrderTime = orderTime2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetOrderTime() int64 {
	return this.OrderTime
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetShopName(shopName2 string) {
	this.ShopName = shopName2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetShopName() string {
	return this.ShopName
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetUserName(userName2 string) {
	this.UserName = userName2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetUserName() string {
	return this.UserName
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.order.push.beta"
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["actual_amount"] = this.ActualAmount
	txtParams["allowance_amount"] = this.AllowanceAmount
	txtParams["ext"] = this.Ext
	txtParams["meal_plan_no"] = this.MealPlanNo
	txtParams["meal_time"] = this.MealTime
	txtParams["order_details"] = this.OrderDetails
	txtParams["order_full_amount"] = this.OrderFullAmount
	txtParams["order_id"] = this.OrderId
	txtParams["order_time"] = this.OrderTime
	txtParams["shop_id"] = this.ShopId
	txtParams["shop_name"] = this.ShopName
	txtParams["user_name"] = this.UserName
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActualAmount, "actualAmount"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowanceOrderPushBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowanceOrderPushBetaResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
