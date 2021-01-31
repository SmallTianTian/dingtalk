package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowanceOrderPushRequest() *OapiCateringAllowanceOrderPushRequest {
	return &OapiCateringAllowanceOrderPushRequest{}
}

type OapiCateringAllowanceOrderPushRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowanceOrderPushResponse
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

func (this *OapiCateringAllowanceOrderPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowanceOrderPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowanceOrderPushRequest) SetActualAmount(actualAmount2 int64) {
	this.ActualAmount = actualAmount2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetActualAmount() int64 {
	return this.ActualAmount
}
func (this *OapiCateringAllowanceOrderPushRequest) SetAllowanceAmount(allowanceAmount2 int64) {
	this.AllowanceAmount = allowanceAmount2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetAllowanceAmount() int64 {
	return this.AllowanceAmount
}
func (this *OapiCateringAllowanceOrderPushRequest) SetExt(ext2 string) {
	this.Ext = ext2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetExt() string {
	return this.Ext
}
func (this *OapiCateringAllowanceOrderPushRequest) SetMealPlanNo(mealPlanNo2 string) {
	this.MealPlanNo = mealPlanNo2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetMealPlanNo() string {
	return this.MealPlanNo
}
func (this *OapiCateringAllowanceOrderPushRequest) SetMealTime(mealTime2 int64) {
	this.MealTime = mealTime2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetMealTime() int64 {
	return this.MealTime
}
func (this *OapiCateringAllowanceOrderPushRequest) SetOrderDetails(orderDetails2 string) {
	this.OrderDetails = orderDetails2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetOrderDetails() string {
	return this.OrderDetails
}
func (this *OapiCateringAllowanceOrderPushRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringAllowanceOrderPushRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringAllowanceOrderPushRequest) SetOrderTime(orderTime2 int64) {
	this.OrderTime = orderTime2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetOrderTime() int64 {
	return this.OrderTime
}
func (this *OapiCateringAllowanceOrderPushRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringAllowanceOrderPushRequest) SetShopName(shopName2 string) {
	this.ShopName = shopName2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetShopName() string {
	return this.ShopName
}
func (this *OapiCateringAllowanceOrderPushRequest) SetUserName(userName2 string) {
	this.UserName = userName2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetUserName() string {
	return this.UserName
}
func (this *OapiCateringAllowanceOrderPushRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowanceOrderPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.order.push"
}
func (this *OapiCateringAllowanceOrderPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowanceOrderPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowanceOrderPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowanceOrderPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowanceOrderPushRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiCateringAllowanceOrderPushRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActualAmount, "actualAmount"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowanceOrderPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowanceOrderPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowanceOrderPushResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
}
