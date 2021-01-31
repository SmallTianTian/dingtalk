package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiNewmanufacturerOrderGetRequest() *OapiNewmanufacturerOrderGetRequest {
	return &OapiNewmanufacturerOrderGetRequest{}
}

type OapiNewmanufacturerOrderGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiNewmanufacturerOrderGetResponse
	Number          string
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiNewmanufacturerOrderGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiNewmanufacturerOrderGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiNewmanufacturerOrderGetRequest) SetNumber(number2 string) {
	this.Number = number2
}
func (this *OapiNewmanufacturerOrderGetRequest) GetNumber() string {
	return this.Number
}
func (this *OapiNewmanufacturerOrderGetRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiNewmanufacturerOrderGetRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiNewmanufacturerOrderGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.newmanufacturer.order.get"
}
func (this *OapiNewmanufacturerOrderGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiNewmanufacturerOrderGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiNewmanufacturerOrderGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiNewmanufacturerOrderGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiNewmanufacturerOrderGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["number"] = this.Number
	txtParams["tenant_id"] = this.TenantId
	return txtParams
}
func (this *OapiNewmanufacturerOrderGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiNewmanufacturerOrderGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiNewmanufacturerOrderGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiNewmanufacturerOrderGetResponse struct {
	taobao.TaobaoResponse
	Model   ProductOrderDto `json:"model,omitempty"`
	Success string          `json:"success,omitempty"`
}
type ProductOrderDto struct {
	Number string `json:"number,omitempty"`
	Status string `json:"status,omitempty"`
}
