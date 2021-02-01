package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringOpenOrderPushRequest() *OapiCateringOpenOrderPushRequest {
	return &OapiCateringOpenOrderPushRequest{}
}

type OapiCateringOpenOrderPushRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringOpenOrderPushResponse
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

func (this *OapiCateringOpenOrderPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringOpenOrderPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringOpenOrderPushRequest) SetActualAmount(actualAmount2 int64) {
	this.ActualAmount = actualAmount2
}
func (this *OapiCateringOpenOrderPushRequest) GetActualAmount() int64 {
	return this.ActualAmount
}
func (this *OapiCateringOpenOrderPushRequest) SetAllowanceAmount(allowanceAmount2 int64) {
	this.AllowanceAmount = allowanceAmount2
}
func (this *OapiCateringOpenOrderPushRequest) GetAllowanceAmount() int64 {
	return this.AllowanceAmount
}
func (this *OapiCateringOpenOrderPushRequest) SetExt(ext2 string) {
	this.Ext = ext2
}
func (this *OapiCateringOpenOrderPushRequest) GetExt() string {
	return this.Ext
}
func (this *OapiCateringOpenOrderPushRequest) SetMealPlanNo(mealPlanNo2 string) {
	this.MealPlanNo = mealPlanNo2
}
func (this *OapiCateringOpenOrderPushRequest) GetMealPlanNo() string {
	return this.MealPlanNo
}
func (this *OapiCateringOpenOrderPushRequest) SetMealTime(mealTime2 int64) {
	this.MealTime = mealTime2
}
func (this *OapiCateringOpenOrderPushRequest) GetMealTime() int64 {
	return this.MealTime
}
func (this *OapiCateringOpenOrderPushRequest) SetOrderDetails(orderDetails2 string) {
	this.OrderDetails = orderDetails2
}
func (this *OapiCateringOpenOrderPushRequest) GetOrderDetails() string {
	return this.OrderDetails
}
func (this *OapiCateringOpenOrderPushRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringOpenOrderPushRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringOpenOrderPushRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringOpenOrderPushRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringOpenOrderPushRequest) SetOrderTime(orderTime2 int64) {
	this.OrderTime = orderTime2
}
func (this *OapiCateringOpenOrderPushRequest) GetOrderTime() int64 {
	return this.OrderTime
}
func (this *OapiCateringOpenOrderPushRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringOpenOrderPushRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringOpenOrderPushRequest) SetShopName(shopName2 string) {
	this.ShopName = shopName2
}
func (this *OapiCateringOpenOrderPushRequest) GetShopName() string {
	return this.ShopName
}
func (this *OapiCateringOpenOrderPushRequest) SetUserName(userName2 string) {
	this.UserName = userName2
}
func (this *OapiCateringOpenOrderPushRequest) GetUserName() string {
	return this.UserName
}
func (this *OapiCateringOpenOrderPushRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringOpenOrderPushRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringOpenOrderPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.open.order.push"
}
func (this *OapiCateringOpenOrderPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringOpenOrderPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringOpenOrderPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringOpenOrderPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringOpenOrderPushRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiCateringOpenOrderPushRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActualAmount, "actualAmount"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringOpenOrderPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringOpenOrderPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringOpenOrderPushResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
