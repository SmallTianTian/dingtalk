package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowanceCapacityGetBetaRequest() *OapiCateringAllowanceCapacityGetBetaRequest {
	return &OapiCateringAllowanceCapacityGetBetaRequest{}
}

type OapiCateringAllowanceCapacityGetBetaRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowanceCapacityGetBetaResponse
	MeaTime         int64
	OrderFullAmount int64
	Orderid         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) SetMeaTime(meaTime2 int64) {
	this.MeaTime = meaTime2
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetMeaTime() int64 {
	return this.MeaTime
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) SetOrderid(orderid2 string) {
	this.Orderid = orderid2
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetOrderid() string {
	return this.Orderid
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.capacity.get.beta"
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mea_time"] = this.MeaTime
	txtParams["order_full_amount"] = this.OrderFullAmount
	txtParams["orderid"] = this.Orderid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MeaTime, "meaTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowanceCapacityGetBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowanceCapacityGetBetaResponse struct {
	taobao.TaobaoResponse
	Result CateringOpenGetDeductCapacityResponse `json:"result,omitempty"`
}
type CateringOpenGetDeductCapacityResponse struct {
	DeductCapacity int64 `json:"deduct_capacity,omitempty"`
}
