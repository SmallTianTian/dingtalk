package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowanceCapacityGetRequest() *OapiCateringAllowanceCapacityGetRequest {
	return &OapiCateringAllowanceCapacityGetRequest{}
}

type OapiCateringAllowanceCapacityGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowanceCapacityGetResponse
	MeaTime         int64
	OrderFullAmount int64
	Orderid         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringAllowanceCapacityGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowanceCapacityGetRequest) SetMeaTime(meaTime2 int64) {
	this.MeaTime = meaTime2
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetMeaTime() int64 {
	return this.MeaTime
}
func (this *OapiCateringAllowanceCapacityGetRequest) SetOrderFullAmount(orderFullAmount2 int64) {
	this.OrderFullAmount = orderFullAmount2
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetOrderFullAmount() int64 {
	return this.OrderFullAmount
}
func (this *OapiCateringAllowanceCapacityGetRequest) SetOrderid(orderid2 string) {
	this.Orderid = orderid2
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetOrderid() string {
	return this.Orderid
}
func (this *OapiCateringAllowanceCapacityGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.capacity.get"
}
func (this *OapiCateringAllowanceCapacityGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowanceCapacityGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowanceCapacityGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mea_time"] = this.MeaTime
	txtParams["order_full_amount"] = this.OrderFullAmount
	txtParams["orderid"] = this.Orderid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringAllowanceCapacityGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MeaTime, "meaTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowanceCapacityGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowanceCapacityGetResponse struct {
	taobao.TaobaoResponse
	Result CateringOpenGetDeductCapacityResponse `json:"result,omitempty"`
}
